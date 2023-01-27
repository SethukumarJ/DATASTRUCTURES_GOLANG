package main

import "fmt"

func main() {

	var series int

	series = fibonacci(7)
	fmt.Println(series)
}

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
