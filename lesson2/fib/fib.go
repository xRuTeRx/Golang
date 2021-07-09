package fib

func Findf(x int) int {
	if x <= 1 {
		return x
	} else {
		return Findf(x-1) + Findf(x-2)
	}

}
