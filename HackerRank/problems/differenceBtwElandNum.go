package main

import "fmt"

func main() {

	nums := []int{1, 23}
	println(differenceBtweenSumOfElementsAndSumOfDigitsOfElements(nums))

}

// differenceBtween sum of the elements and the sum of digits of the elements in a slice

func differenceBtweenSumOfElementsAndSumOfDigitsOfElements(nums []int) int {

	sumOfElements := 0
	sumOfDigitsOfElements := 0

	for _, num := range nums {
		sumOfElements += num
		sumOfDigitsOfElements += sumOfDigits(num)
	}

	fmt.Println("sumOfElements: ", sumOfElements)
	fmt.Println("sumOfDigitsOfElements: ", sumOfDigitsOfElements);

	return sumOfElements - sumOfDigitsOfElements
}

func sumOfDigits(num int) int {

	sum := 0

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}
