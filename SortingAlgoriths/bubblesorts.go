package main

import "fmt"

func bubbleSort(array []int) {

	for i := 0; i < len(array); i++ {

		for j := 0; j < len(array)-(i+1); j++ {
		
			if array[j] >= array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}

		}
	}

}

func main() {

	arr := []int{4, 3, 45, 5, 3, 5, 3, 25}
	bubbleSort(arr)

	fmt.Println(arr)
}
