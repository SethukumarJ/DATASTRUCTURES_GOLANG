package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}

	index := binarySearch(arr, 0, 4, 2)

	fmt.Println(index)
}

func binarySearch(arr []int, l, r, x int) int {
	if r >= l {
		mid := l + (r-l)/2
		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		}
		return binarySearch(arr, mid+1, r, x)
	}
	return -1
}
