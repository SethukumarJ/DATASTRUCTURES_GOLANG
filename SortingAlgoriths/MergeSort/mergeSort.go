package main


func mergeSortHelper(array []int, firstIdx int,lastIdx int) {

	midIdx := (firstIdx + lastIdx)/2
	

	
	if firstIdx == lastIdx {

		return
	}

	mergeSortHelper(array , firstIdx, midIdx)
	mergeSortHelper(array, midIdx+1, lastIdx)


}





func mergeSort(array []int) {

}