package main 

import (

	"fmt"

)

type MinHeap struct {
	 arr []int
}

func(h *MinHeap) insert(data int) {

	h.arr = append(h.arr, data)
	h.insertHelper(len(h.arr)-1)

}


func (h *MinHeap) insertHelper(position int) {

	if h.arr[position] < h.arr[(position-1)/2] {
		swap(h.arr,position,(position-1)/2)
		h.insertHelper((position-1)/2)
	}
}

func (h *MinHeap) rootDelete() {

	   h.arr[0] = h.arr[len(h.arr)-1]
       h.arr = h.arr[:len(h.arr)-1]
       h.rootDeleteHelper(0)

    }
   
 
func (h *MinHeap) rootDeleteHelper(position int) {

	 if h.arr[position] > h.arr[h.minValue(position)] {
        swap(h.arr,position,h.minValue(position))
        h.rootDeleteHelper(h.minValue(position));
	}
}

func (h *MinHeap) minValue(position int) int{
       
	if(h.arr[2*position+1] > h.arr[2*position+2]){

		return 2*position+2;

	} else {

	return 2*position+1;

	}
}

func (h  *MinHeap) display() {

	for i:= 0; i< len(h.arr); i++ {

		fmt.Print(h.arr[i]," ")
	}
}

func(t *MinHeap) BuildHeap(array []int) {
	t.arr = array
	size := len(array)
	for i := parent(size-1); i>=0; i--{
		t.shiftDown(i)
	}
}



func (h *MinHeap) shiftDown(currentIdx int) {

	endIdx := len(h.arr)-1
	leftIdx := leftChild(currentIdx)

	for leftIdx <= endIdx {

		rightIdx := rightChild(currentIdx)
		var idxToShift int  

		if rightIdx <= endIdx && h.arr[rightIdx] < h.arr[leftIdx] {
			idxToShift = rightIdx
		} else {
			idxToShift = leftIdx
		}

		if h.arr[idxToShift] < h.arr[currentIdx] {
			swap(h.arr, idxToShift,currentIdx)
			currentIdx = idxToShift
			leftIdx = leftChild(currentIdx)
		} else {
			return
		}
	}
}




func main () {

heap := MinHeap{}

array := []int{10,49,3,1,5,7,12,5}

heap.BuildHeap(array)
heap.insert(9)
heap.rootDelete()
heap.display()

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
// func (h *MinHeap) shiftUp(currentIdx int) {
// 	parentIdx := parent(currentIdx)

// 	for currentIdx >0 && h.arr[parentIdx] > h.arr[currentIdx] {
// 		swap(h.arr,currentIdx,parentIdx)
// 		currentIdx = parentIdx
// 		parentIdx = parent(currentIdx  )
// 	}
// }
 