package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "{")
	fmt.Fprintln(w, "	\"Host\": "+r.Host+",")
	fmt.Fprintln(w, "	\"UserAgent\": "+r.UserAgent()+",")
	fmt.Fprintln(w, "	\"RequestURL\": "+r.URL.Path+",")
	fmt.Fprintln(w, "	\"Headers\": {")
	for k, v := range r.Header {
		fmt.Fprint(w, "		\""+k+"\" : ")
		fmt.Fprint(w, v)
		fmt.Fprintln(w, ",")
	}
	fmt.Fprintln(w, "	} ")
	fmt.Fprintln(w, "}")

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
