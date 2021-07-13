package main

import "fmt"

func max(s []string) string {
	max := ""
	for _, vs := range s {
		if len(vs) > len(max) {
			max = vs
		}
	}
	return max
}
func main() {
	s := []string{"first", "second", "third", "fours"}
	fmt.Println(max(s))
}
