package main 

import (

	"fmt"

)

type MinHeap struct {

	 arr [100]int

}

var size int = -1


func(h *MinHeap) insert(data int) {

	h.arr[size+1] = data
	size++
	h.insertHelper(size)

}




func (h *MinHeap) insertHelper(position int) {

	if h.arr[position] < h.arr[(position-1)/2] {

		swap := h.arr[position]
		h.arr[(position-1)/2] = h.arr[position]
		h.arr[position] = swap
		h.insertHelper((position-1)/2)

	}

}




func main () {



}