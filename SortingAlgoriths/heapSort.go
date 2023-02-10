package main

import "fmt"

func HeapSort(array []int, endIdx int) []int{
	if endIdx <= 0 {
		return array
	}
	sorted := Heapify(array,endIdx)
	swap(sorted,endIdx,0)
	return HeapSort(sorted, endIdx-1)

}


func Heapify(array []int, endIdx int) []int{
	

	for i:= parent(endIdx); i>=0; i--{
		ShiftDown(array,endIdx,i)
	}

	return array
}	

func ShiftDown(array []int, endIdx , currentIdx int) []int{

	leftIdx := leftChild(currentIdx)

	for leftIdx <= endIdx {

			rightIdx := rightChild(currentIdx)
			var idxToShift int
			if rightIdx <= endIdx && array[rightIdx]>array[leftIdx]{
				idxToShift = rightIdx
			} else {
				idxToShift = leftIdx
			}


			if array[idxToShift] > array[currentIdx] {

				swap(array,idxToShift,currentIdx)
				currentIdx = idxToShift
				leftIdx = leftChild(currentIdx)
			} else {
				return array
			}
	}

	return array
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