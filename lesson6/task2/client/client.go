package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// form key for new message
const keyString = "InputString"

func main() {
	// server port number
	const port = 8080
	// server address
	address := fmt.Sprintf("http://127.0.0.1:%d", port)

	fmt.Printf("Connecting to server %s\n", address)
	for {

		r := bufio.NewReader(os.Stdin)
		fmt.Print("Enter value: ")
		inputValue, _ := r.ReadString('\n')
		inputValue = strings.TrimSpace(inputValue)
		if inputValue == "exit" {
			break
		}

		// post new message
		resp, err := http.PostForm(address, url.Values{keyString: {inputValue}})
		if err != nil {
			log.Fatal(err)
		}

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Output string : " + string(respBody))
		fmt.Println("------------------------------------")
	}
	fmt.Println("DONE!")
}
