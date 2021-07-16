package main

import "fmt"

func reverse(s []int64) []int64 {
	rs := make([]int64, len(s))
	for i, vs := range s {
		rs[len(s)-i-1] = vs
	}
	return rs
}
func main() {
	s := []int64{1, 2, 5, 2, 1, 5, 2}
	fmt.Println(reverse(s))
}
