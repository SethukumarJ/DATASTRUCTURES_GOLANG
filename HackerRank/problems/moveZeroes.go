package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j:= i+1; j< len(nums); j++{
				if nums[j] != 0 {
					nums[i],nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}
