package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	ID          int    `json:"ID"`
	City        string `json:"City"`
	RequestTime string `json:"RequestTime"`
	Temperature string `json:"Temperature"`
}

const (
	host = "localhost"
	port = 8080
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Input 1 if u want list requests")
		fmt.Println("Input 2 if u want view temperature in city")
		fmt.Println("Input 3 if u want an exit")
		fmt.Print("Enter value: ")
		choise, _ := r.ReadString('\n')
		choise = strings.TrimSpace(choise)
		switch choise {
		case "1":
			url := fmt.Sprintf("http://%s:%d/listRequests", host, port)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			var result []Request
			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(result)

		case "2":
			//r := bufio.NewReader(os.Stdin)
			fmt.Print("Enter city name: ")
			choise, _ := r.ReadString('\n')
			choise = strings.TrimSpace(choise)
			url := fmt.Sprintf("http://%s:%d/weather?city=%s", host, port, choise)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				var result string
				err = json.NewDecoder(resp.Body).Decode(&result)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Printf("%s\n", result)
			} else {
				var result float64
				err = json.NewDecoder(resp.Body).Decode(&result)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Printf("temperature in %s : %.2f\n", choise, result)
			}
		case "3":
			os.Exit(0)
		default:
			fmt.Println("wrong choise!")
		}
		fmt.Println("------------------------------------")

	}

}
