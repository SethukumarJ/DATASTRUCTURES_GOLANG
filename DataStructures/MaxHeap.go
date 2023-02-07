package main 

import (

	"fmt"

)

type MaxHeap struct {

	 arr []int

}

func(h *MaxHeap) insert(data int) {

	h.arr = append(h.arr, data)
	h.insertHelper(len(h.arr)-1)

}


func (h *MaxHeap) insertHelper(position int) {

	if h.arr[position] > h.arr[(position-1)/2] {
		swap(h.arr,position,(position-1)/2)
		h.insertHelper((position-1)/2)
	}
}

func (h *MaxHeap) rootDelete() {

	   h.arr[0] = h.arr[len(h.arr)-1]
       h.arr = h.arr[:len(h.arr)-1]
       h.rootDeleteHelper(0)

    }
   
 
func (h *MaxHeap) rootDeleteHelper(position int) {

	 if h.arr[position] < h.arr[h.minValue(position)] {
        swap(h.arr,position,h.minValue(position))
        h.rootDeleteHelper(h.minValue(position));
	}
}

func (h *MaxHeap) minValue(position int) int{
       
	if(h.arr[2*position+1] > h.arr[2*position+2]){

		return 2*position+2;

	} else {

	return 2*position+1;

	}
}

func (h  *MaxHeap) display() {

	for i:= 0; i< len(h.arr); i++ {

		fmt.Print(h.arr[i]," ")
	}
}

func (h *MaxHeap) BuildHeap(array []int ) {
	h.arr = array
	for i:= parent(len(h.arr)-1); i>=0; i--{

		h.shiftDown(i)
	}
}





func (h *MaxHeap) shiftDown(currentIdx int) {

	endIdx := len(h.arr)-1
	leftIdx := leftChild(currentIdx)

	for leftIdx <= endIdx {

		rightIdx := rightChild(currentIdx)
		var idxToShift int 
		if rightIdx <= endIdx && h.arr[rightIdx] > h.arr[leftIdx] {
			idxToShift = rightIdx
			
		} else {
			idxToShift = leftIdx
		}

		if h.arr[idxToShift] > h.arr[currentIdx] {
			swap(h.arr, currentIdx,idxToShift)
		} else {
			return
		}
	}


}

func main () {

heap := MaxHeap{}

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