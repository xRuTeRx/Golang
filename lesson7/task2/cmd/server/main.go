package main

import (
	"log"
	"net/http"

	"github.com/xruterx/golang/lesson7/task2/handlers"
)

func main() {
	addr := "localhost:8081"
	endpoint := "/rpc"
	http.HandleFunc(endpoint, handlers.Handle)
	log.Printf("Started JSON RPC server on %s%s", addr, endpoint)
	log.Fatal(http.ListenAndServe(addr, nil))
}
