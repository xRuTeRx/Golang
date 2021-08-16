package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

const (
	id_key      = "3b2ea0ea927a17e540f2330bcafdaba3" //my id_key for api.openweathermap.org
	serverPort  = 8080
	dbType      = "postgres"
	dbPort      = 5432
	table       = "requests"
	userDB      = "root"
	pwdDB       = "root"
	hostDB      = "localhost"
	weatherHost = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s" //with get
)

type SS struct {
	NoErr bool        `json:"noErr"`
	Mes   interface{} `json:"mes"`
}

type Request struct {
	ID          int    `db:"id"`
	City        string `db:"city"`
	RequestTime string `db:"time"`
	Temperature string `db:"temperature"`
}

func connString(dbType, user, pwd, dbHost string, port int, dbName string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", dbType, user, pwd, dbHost, port, dbName)
}

func rowsToRequest(rows *sql.Rows) (request []Request, err error) {
	Requests := make([]Request, 0)
	for rows.Next() {
		u := &Request{}
		if err = rows.Scan(&u.ID, &u.City, &u.RequestTime, &u.Temperature); err != nil {
			return nil, errors.Wrap(err, "failed to list all requests (scan)")
		}
		Requests = append(Requests, *u)
	}
	return Requests, nil
}

func (db *DB) ListAll() ([]Request, error) {
	q := "select * from requests"
	rows, err := db.Conn.Query(q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list all requests")
	}
	return rowsToRequest(rows)
}
func (db *DB) AddRequest(cityName string, time string, temperature string) (int64, error) {
	q := "INSERT INTO requests (city, time, temperature) VALUES ($1, $2, $3);"
	res, err := db.Conn.Exec(q, cityName, time, temperature)
	if err != nil {
		return 0, errors.Wrap(err, "failed to add request")
	}
	insertedId, _ := res.LastInsertId()
	return insertedId, nil
}

func (db *DB) UpdateRequestTemp(id int64, temperature string) error {
	q := "UPDATE  requests SET  temperature=$1 WHERE ID=$2 ;"
	if _, err := db.Conn.Exec(q, temperature, id); err != nil {
		return errors.Wrap(err, "failed to add request")
	}
	return nil
}

func ConnToDB() (*DB, error) {
	connStr := connString(dbType, userDB, pwdDB, hostDB, dbPort, table)
	c, err := sql.Open(dbType, connStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to DB")
	}
	db := &DB{
		Conn: c,
	}
	return db, nil

}
func writeToWeb(w http.ResponseWriter, mes interface{}, cod bool) {
	s := &SS{
		Mes:   mes,
		NoErr: cod,
	}
	b, err := json.Marshal(s)
	if err != nil {
		return
	}
	w.Write(b)
}

func handlerWeather(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/weather" {
		writeToWeb(w, "Page not found", false)
		return
	}
	switch r.Method {
	case http.MethodGet:
		db, err := ConnToDB()
		if err != nil {
			writeToWeb(w, err, false)
			return
		}
		defer db.Conn.Close()
		cityName := r.FormValue("city")
		insertedID, err := db.AddRequest(cityName, time.Now().String(), "")
		if err != nil {
			writeToWeb(w, err, false)
			return
		}

		url := fmt.Sprintf(weatherHost, cityName, id_key)
		resp, err := http.Get(url)
		if err != nil {
			writeToWeb(w, err, false)
			return
		}

		defer resp.Body.Close()

		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		cod, _ := json.Marshal(result["cod"])
		if string(cod) != "200" {
			res := result["message"]
			writeToWeb(w, res, false)
		} else {
			res, err := json.Marshal(result["main"])
			if err != nil {
				writeToWeb(w, err, false)
				return
			}
			err = json.Unmarshal(res, &result)
			if err != nil {
				writeToWeb(w, err, false)
				return
			}
			res, err = json.Marshal(result["temp"])
			if err != nil {
				writeToWeb(w, err, false)
				return
			}

			err = db.UpdateRequestTemp(insertedID, string(res))
			if err != nil {
				writeToWeb(w, err, false)
				return
			}
			writeToWeb(w, string(res), true)
		}
	default:
		writeToWeb(w, "Sorry, only GET methods are supported.", false)
	}
}

func handlerListRequests(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/listRequests" {

		writeToWeb(w, "Page not found", false)
		return
	}
	switch r.Method {
	case http.MethodGet:
		db, err := ConnToDB()
		defer db.Conn.Close()
		if err != nil {
			writeToWeb(w, err, false)
			return
		}

		res, err := db.ListAll()
		if err != nil {
			writeToWeb(w, err, false)
			return
		}

		writeToWeb(w, res, true)

	default:
		writeToWeb(w, "Sorry, only GET methods are supported.", false)
	}
}

func main() {
	http.HandleFunc("/weather", handlerWeather)
	http.HandleFunc("/listRequests", handlerListRequests)
	// start server without ending
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil))
}
