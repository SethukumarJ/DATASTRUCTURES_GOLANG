package main

import "fmt"

func HeapSort(array []int, endIdx int) []int {
	if endIdx < 0 {
		return array
	}

	array = Heapify(array, endIdx)
	swap(array, 0, endIdx)
	return HeapSort(array, endIdx-1)
}

func Heapify(array []int, endIdx int) []int {
	for i := parent(endIdx); i >= 0; i-- {
		shiftDown(i, array, endIdx)
	}

	return array
}

func shiftDown(currentIdx int, arr []int, endIdx int) {
	leftIdx := leftChild(currentIdx)

	for leftIdx <= endIdx {
		rightIdx := rightChild(currentIdx)
		var idxToShift int
		if rightIdx <= endIdx && arr[rightIdx] > arr[leftIdx] {
			idxToShift = rightIdx
		} else {
			idxToShift = leftIdx
		}

		if arr[idxToShift] > arr[currentIdx] {
			swap(arr, currentIdx, idxToShift)
			currentIdx = idxToShift
			leftIdx = leftChild(currentIdx)
		} else {
			break
		}
	}
}

func main() {
	array := []int{5, 6, 3, 6, 3, 2, 17}

	sortedArray := HeapSort(array, len(array)-1)

	fmt.Println("Sorted array:", sortedArray)
}




func swap(arr []int, i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}

func leftChild(i int) int{
	return (i*2) +1
}
func rightChild(i int) int{
	return (i*2) +2
}
func parent(i int) int{
	return(i-1)/2
}