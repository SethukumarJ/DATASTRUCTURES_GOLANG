package main

import "fmt"

func main() {

	facorials := factorial(6)

	fmt.Println(facorials)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
