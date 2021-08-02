package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Tokens struct {
	Token     string `json:"token"`
	CreatedAt string `json:"createdAt"`
	ExpiredAt string `json:"expireAt"`
}

var TokenSlice []Tokens

func handler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/saveCoockie" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodPost:
		body := r.Body
		defer body.Close()
		bodyData, err := io.ReadAll(body)
		if err != nil {
			log.Print(err.Error())
			return
		}
		req := &Tokens{}

		if err = json.Unmarshal(bodyData, req); err != nil {
			log.Print(err.Error())
			return
		}
		TokenSlice = append(TokenSlice, *req)
		fmt.Println(TokenSlice)
		fmt.Println("------------------------------")

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	// server port number
	const port = 8081

	fmt.Printf("Launching server on port: %d \n\n", port)

	// set handler for route '/'
	http.HandleFunc("/saveCoockie", handler)
	// start server without ending
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
