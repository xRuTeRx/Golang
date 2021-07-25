package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		body, _ := ioutil.ReadFile("../form/form.html")
		fmt.Fprint(w, string(body))

	case "POST":
		cookie := http.Cookie{
			Name:  "token",
			Value: r.PostFormValue("name") + ":" + r.PostFormValue("address"),
		}
		http.SetCookie(w, &cookie)
		body, _ := ioutil.ReadFile("../form/form.html")
		fmt.Fprint(w, string(body))

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
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
