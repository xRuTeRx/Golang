package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/xruterx/golang/last_task/db"
	f "github.com/xruterx/golang/last_task/server/handler/function"
)

func HandlerWeather(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/weather" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		db, err := db.ConnToDB()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Conn.Close()
		cityName := r.FormValue("city")
		insertedID, err := db.AddRequest(cityName, time.Now().String(), "")
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server database error", http.StatusInternalServerError)
			return
		}
		res, err := f.WeatherFromOpenweathermap(w, cityName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = db.UpdateRequestTemp(insertedID, string(res))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server database error", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Sorry, only GET methods are supported.", http.StatusNotImplemented)
	}
}

func HandlerListRequests(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/listRequests" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		db, err := db.ConnToDB()
		defer db.Conn.Close()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server database error", http.StatusInternalServerError)
			return
		}

		res, err := db.ListAll()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "internal server database error", http.StatusInternalServerError)
			return
		}

		message, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(message)

	default:
		http.Error(w, "Sorry, only GET methods are supported.", http.StatusNotImplemented)
	}
}
