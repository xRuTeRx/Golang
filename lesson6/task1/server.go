package main

import (
	"fmt"
	"log"
	"net/http"
)

type OutStruct struct {
	Host       string
	UserAgent  string
	RequestUrl string
	Headers    http.Header
}

func handler(w http.ResponseWriter, r *http.Request) {
	var outS OutStruct
	outS.Host = r.Host
	outS.UserAgent = r.UserAgent()
	outS.RequestUrl = r.URL.Path
	outS.Headers = r.Header
	fmt.Fprint(w, outS)

}

func main() {
	// server port number
	const port = 8080

	fmt.Printf("Launching server on port: %d \n\n", port)

	// set handler for route '/'
	http.HandleFunc("/", handler)
	// start server without ending
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
