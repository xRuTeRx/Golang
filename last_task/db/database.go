package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/pkg/errors"
	st "github.com/xruterx/golang/last_task/db/structure"
)

type DB struct {
	Conn *sql.DB
}

func connString(dbType, user, pwd, dbHost string, port string, dbName string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbType, user, pwd, dbHost, port, dbName)
}

func rowsToRequest(rows *sql.Rows) (request []st.Request, err error) {
	Requests := make([]st.Request, 0)
	for rows.Next() {
		u := &st.Request{}
		if err = rows.Scan(&u.Id, &u.City, &u.RequestTime, &u.Temperature); err != nil {
			return nil, errors.Wrap(err, "failed to list all requests (scan)")
		}
		Requests = append(Requests, *u)
	}
	return Requests, nil
}

func (db *DB) ListAll() ([]st.Request, error) {
	q := "select * from requests"
	rows, err := db.Conn.Query(q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list all requests")
	}
	return rowsToRequest(rows)
}
func (db *DB) AddRequest(cityName string, time string, temperature string) (int64, error) {
	q := "INSERT INTO \"requests\" (\"City\", \"RequestTime\", \"Temperature\") VALUES ($1, $2, $3) RETURNING \"Id\";"
	var insertedId int64
	err := db.Conn.QueryRow(q, cityName, time, temperature).Scan(&insertedId)
	if err != nil {
		return 0, errors.Wrap(err, "failed to add request")
	}

	return insertedId, nil
}

func (db *DB) UpdateRequestTemp(id int64, temperature string) error {
	q := "UPDATE  requests SET  \"Temperature\"=$1 WHERE \"Id\"=$2 ;"
	if _, err := db.Conn.Exec(q, temperature, id); err != nil {
		return errors.Wrap(err, "failed to update request")
	}
	return nil
}

func ConnToDB() (*DB, error) {
	connStr := connString(os.Getenv("DB_TYPE"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_TABLE"))
	c, err := sql.Open(os.Getenv("DB_TYPE"), connStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to DB")
	}
	db := &DB{
		Conn: c,
	}
	return db, nil

}
