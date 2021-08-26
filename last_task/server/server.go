package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/xruterx/golang/last_task/server/handler"
)

func main() {
	err := godotenv.Load("db.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fmt.Println("start")
	http.HandleFunc("/weather", handler.HandlerWeather)
	http.HandleFunc("/listRequests", handler.HandlerListRequests)
	// start server without ending
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), nil))
}
