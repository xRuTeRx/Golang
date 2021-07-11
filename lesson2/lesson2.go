package main

import (
	"fmt"
	"lesson2/fib"
)

func main() {
	defer fmt.Println("Done!")
	fmt.Println("20 Fibanachi mbers")
	fib.OutFibanacciLine(20)

}
