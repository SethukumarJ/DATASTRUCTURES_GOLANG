package main

import "fmt"

func main() {

	input := []int{7, 2, 5, 3, 6, 10}
	fmt.Println("input-", input)
	var slice int
	slice = maxProfit(input)

	fmt.Println("output-", slice)

}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}

	return profit

}
