package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Tokens struct {
	Token     string `json:"token"`
	CreatedAt string `json:"createdAt"`
	ExpiredAt string `json:"expireAt"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		body, _ := ioutil.ReadFile("form/form.html")
		fmt.Fprint(w, string(body))

	case http.MethodPost:
		cookie := http.Cookie{
			Name:    "token",
			Value:   r.PostFormValue("name") + ":" + r.PostFormValue("address"),
			Expires: time.Now().Add(10 * 24 * time.Hour),
		}
		http.SetCookie(w, &cookie)

		// request to anather server /saveCoockie
		postBody, _ := json.Marshal(Tokens{
			Token:     cookie.Value,
			CreatedAt: time.Now().String(),
			ExpiredAt: cookie.Expires.String(),
		})
		fmt.Println(string(postBody))
		respBuffer := bytes.NewBuffer(postBody)

		_, err := http.Post("http://savecoockie:8081/saveCoockie", "application/json", respBuffer)
		if err != nil {
			log.Print(err.Error())
			return
		}
		body, _ := ioutil.ReadFile("form/form.html")
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
