package main

import "fmt"

func main() {
	input := []int{1, 2, 3}
	fmt.Println("input-",input)
	slice := []int{}
	slice = PlusOne(input)
	
	fmt.Println("output-",slice)
}

func PlusOne(x []int) []int {

	for i := len(x) - 1; i >= 0; i-- {

		x[i]++ 
		if x[i]%10 != 0 {
			return x
		} else {
			x[i] = 0
		}

	}

	return append([]int{1}, x...)
}
 