package fib

import "fmt"

func OutFibanacciLine(x int) {
	fmt.Print("1 1")
	var a, b, c int = 1, 1, 2

	for i := 3; i <= x; i++ {
		c = a + b
		a, b = b, c
		fmt.Print(" ", c)
	}
	fmt.Println("")
}
