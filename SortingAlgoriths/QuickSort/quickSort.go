package main

import "fmt"

func QuickSort(array []int, startIdx int, endIdx int) {
	fmt.Println("array before sorting:", array)
	QuickSortHelper(array, startIdx, endIdx)

	fmt.Println("sorted array :", array)

}

func QuickSortHelper(array []int, startIdx int, endIdx int) {

	if endIdx <= startIdx {
		return
	}

	pivotIdx := startIdx
	leftIdx := startIdx + 1
	rightIdx := endIdx

	for leftIdx <= rightIdx {

		if array[leftIdx] > array[pivotIdx] && array[rightIdx] < array[pivotIdx] {
			
			swap(array,leftIdx,rightIdx)
			leftIdx++
			rightIdx--
			
		}
		if array[leftIdx] <= array[pivotIdx] {
			leftIdx ++
		} 
		if array[rightIdx] >=array[pivotIdx]  {
			rightIdx--
		}
		
	}

	swap(array,rightIdx,pivotIdx)
	QuickSortHelper(array,startIdx, rightIdx-1)
	QuickSortHelper(array,rightIdx+1,endIdx)

}

func swap(array []int, a int, b int) {

	temp := array[a]
	array[a] = array[b]
	array[b] = temp
}


func main() {

	var array = []int{6,4,5,6,7,54,4,3,34,6}
	QuickSort(array,0,9)

}