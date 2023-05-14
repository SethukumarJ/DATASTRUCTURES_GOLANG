package main

import "fmt"

func main() {

	sortedArray := mergeSortedArrays([]int{10, 12, 13, 200}, []int{2, 6, 10})
	fmt.Println(sortedArray)
}

// mergeSortedArrays merges two sorted arrays into one sorted array
func mergeSortedArrays(nums1 []int, nums2 []int) []int {

	if len(nums1) == 0 {
		return nums2
	}

	if len(nums2) == 0 {
		return nums1
	}

	merged := make([]int, len(nums1)+len(nums2))

	i, j := 0, 0

	for k := 0; k < len(merged); k++ {
		if i >= len(nums1) {
			merged[k] = nums2[j]
			j++
			continue
		}
		if j >= len(nums2) {
			merged[k] = nums1[i]
			i++
			continue
		}

		if nums1[i] < nums2[j] {
			merged[k] = nums1[i]
			i++
			continue
		}
		merged[k] = nums2[j]
		j++
	}

	return merged
}
