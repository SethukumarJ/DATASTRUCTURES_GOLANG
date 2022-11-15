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

func (h *MinHeap) rootDelete() {

	   h.arr[0] = h.arr[size]
       size--
       h.rootDeleteHelper(0)

    }
   
 
func (h *MinHeap) rootDeleteHelper(position int) {

	var swap int
        if h.arr[position] > h.minValue(position) {
            swap =h.arr[position];
            h.arr[position] = h.arr[h.minValue(position)];
            h.arr[h.minValue(position)] = swap;
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

	for i:= 0; i< size; i++ {

		fmt.Print(h.arr[i]," ")
	}
}




func main () {



}