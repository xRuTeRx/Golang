package main

import "fmt"

func main() {
	a := [5]int{1, 5, 2, 4, 3}
	fmt.Println(AverageOfArray(a))
}

func AverageOfArray(a [5]int) float32 {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return float32(sum / len(a))
}
