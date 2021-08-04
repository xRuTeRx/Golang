package main

import (
	"log"
	"net/http"

	"github.com/xruterx/golang/lesson7/task2/handlers"
)

func main() {
	addr := "localhost:8081"
	http.HandleFunc("/rpc", handlers.Handle)
	log.Printf("Started JSON RPC server on %s/rpc", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
