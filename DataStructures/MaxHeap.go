package main 

import (

	"fmt"

)

type MaxHeap struct {

	 arr [100]int

}

var size int = -1


func(h *MaxHeap) insert(data int) {

	h.arr[size+1] = data
	size++
	h.insertHelper(size)

}


func (h *MaxHeap) insertHelper(position int) {

	if h.arr[position] > h.arr[(position-1)/2] {

		swap := h.arr[position]
		h.arr[(position-1)/2] = h.arr[position]
		h.arr[position] = swap
		h.insertHelper((position-1)/2)

	}

}

func (h *MaxHeap) rootDelete() {

	   h.arr[0] = h.arr[size]
       size--
       h.rootDeleteHelper(0)

    }
   
 



func (h *MaxHeap) rootDeleteHelper(position int) {

	var swap int
        if h.arr[position] < h.maxValue(position) {
            swap =h.arr[position];
            h.arr[position] = h.arr[h.maxValue(position)];
            h.arr[h.maxValue(position)] = swap;
            h.rootDeleteHelper(h.maxValue(position));

        }


}

func (h *MaxHeap) maxValue(position int) int{
       
	if(h.arr[2*position+1] < h.arr[2*position+2]){

		return 2*position+2;

	} else {

	return 2*position+1;

	}
}

func (h  *MaxHeap) display() {

	for i:= 0; i< size; i++ {

		fmt.Print(h.arr[i]," ")
	}
}




func main () {



}