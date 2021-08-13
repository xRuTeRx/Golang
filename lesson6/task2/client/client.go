package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// create a dialer
	var d net.Dialer
	// server port number
	const port = 8081

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second)*10)
		defer cancel()

		// connect to server with context
		conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatal(err)
		}
		// call close to connection when we end with our tasks
		defer conn.Close()
		// our msg
		r := bufio.NewReader(os.Stdin)
		fmt.Print("Enter value: ")
		message, _ := r.ReadString('\n')
		message = strings.TrimSpace(message)
		if message == "exit" {
			break
		}

		// message[:len(message) - 1] - removes '\n' for logging
		fmt.Printf("Sending message: %s; to port: %d\n", message, port)

		// create call context that should close when timeout reached

		// send some data to server
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			log.Fatal(err)
		}

		// create buffer and read message from server
		getMessage, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Message recieved: %s\n", getMessage[:len(getMessage)-1])

	}

}
