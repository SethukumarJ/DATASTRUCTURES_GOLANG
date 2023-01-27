//Problem using Binary search.
// Find the index of specific character in slice

package main

import "fmt"

var input = []int{1, 2, 3, 4, 5, 6, 7, 8}

func main() {
	number := FindIndex(input, 8)

	fmt.Println(number)
}

func FindIndex(arr []int,number int) int {
	low := 0
    high := len(arr) - 1
    for low <= high {
        mid := (low + high) / 2
        if arr[mid] == number {
            return mid
        } else if arr[mid] < number {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1

}
