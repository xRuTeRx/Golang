package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const keyString = "InputString"

type Request struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Id     int64           `json:"id"`
}
type jsonParams struct {
	Name string `json:"name"`
}

func main() {

	const port = 8081
	address := fmt.Sprintf("http://localhost:%d/rpc", port)
	defer fmt.Println("DONE!")
	for {

		r := bufio.NewReader(os.Stdin)
		fmt.Println("Input 1 if u want add value")
		fmt.Println("Input 2 if u want view list of values")
		fmt.Println("Input 3 if u want an exit")
		fmt.Print("Enter value: ")
		choise, _ := r.ReadString('\n')
		choise = strings.TrimSpace(choise)
		switch choise {
		case "1":
			fmt.Print("Enter what u want to add: ")
			inputValue, _ := r.ReadString('\n')
			inputValue = strings.TrimSpace(inputValue)
			byteInputParams, _ := json.Marshal(jsonParams{
				Name: inputValue,
			})
			postBody, _ := json.Marshal(Request{
				Method: "register",
				Id:     1,
				Params: byteInputParams,
			})
			respBuffer := bytes.NewBuffer(postBody)
			resp, err := http.Post(address, "application/json", respBuffer)
			if err != nil {
				log.Fatalf("An Error Occured %v", err)
			}
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(respBody))
		case "2":
			postBody, _ := json.Marshal(Request{
				Method: "list",
				Id:     2,
				Params: nil,
			})
			respBuffer := bytes.NewBuffer(postBody)
			resp, err := http.Post(address, "application/json", respBuffer)
			if err != nil {
				log.Fatalf("An Error Occured %v", err)
			}
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(respBody))
		case "3":
			os.Exit(0)
		default:
			fmt.Println("wrong choise!")
		}
		fmt.Println("------------------------------------")

	}

}
