package main

import (
	"fmt"
	"sort"
)

func printSorted(m map[int]string) {
	s := make([]int, 0)
	for k := range m {
		s = append(s, k)
	}
	sort.Ints(s)
	for _, v := range s {
		fmt.Print(m[v], " ")
	}
}
func main() {
	var m = map[int]string{2: "2_", 1: "1_", 3: "3_", 10: "10_", 5: "5_"}
	printSorted(m)
}
