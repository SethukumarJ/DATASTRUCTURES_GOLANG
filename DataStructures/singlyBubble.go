package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

func (list *SinglyLinkedList) addNode(data int) {
	newNode := new(Node)
	newNode.Data = data
	newNode.Next = nil

	if list.head == nil {
		list.head = newNode

	}  else {
		list.tail.Next = newNode
	}

	list.tail = newNode

}

func (list *SinglyLinkedList) addNodeBeggining(data int) {
	newNode := new(Node)
	newNode.Data = data
	newNode.Next = nil

	if list.head == nil {
		list.head = newNode

	}  else {
		newNode.Next = list.head
	}

	list.head = newNode

}
func (list *SinglyLinkedList) delete(data int){

	if (list.head != nil && list.head.Data == data) {
		list.head = list.head.Next;
		return
	}
	temp := list.head
	for (temp != nil) {
		if (temp.Next.Data == data) {
			if (temp.Next == list.tail) {
				list.tail = temp;
				list.tail.Next = nil;
			} else {
				temp.Next = temp.Next.Next;

			return
			}
		}

		temp = temp.Next;
	}
  
}

func (list *SinglyLinkedList) display() {
	temp := new(Node)
	if(list.head ==nil) {
		fmt.Println("Empty list!")
	}

	temp = list.head

	for temp !=nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}

}

func (list *SinglyLinkedList) insert(data int,search int, counts int) {
    newNode := new(Node)
	newNode.Data = data
	temp := list.head
	
	count := 0

	
	if search == list.head.Data {
		
		newNode.Next = temp 
		list.head = newNode
	    return
	}

	
	for temp.Next != nil {

		if temp.Next.Data== search {

			count++

			if count == counts {
				newNode.Next = temp.Next
				temp.Next = newNode
				return
			}

		
		}
		temp = temp.Next

	}
}


func (list *SinglyLinkedList) bubbleSort() {

	prev := list.head
	endVariable := prev
	temp := list.head.Next 
	
	for temp != nil && temp != endVariable {
		if temp.Data >= prev.Data {
			temp.Data, prev.Data = prev.Data, temp.Data
			
		}
			prev = temp
			temp = temp.Next

		if temp  == nil || temp == endVariable {
			endVariable = prev
			prev = list.head
			temp = prev.Next
		}
	}

}

func main() {

	list := SinglyLinkedList{}
	list.addNode(3)
	list.addNode(9)
	list.addNode(20)
	list.addNode(1)
	list.addNode(7)
	list.addNode(5)
	list.addNode(2)
	list.addNode(0)
	list.display()
	list.bubbleSort()
	fmt.Println("after buble sort:")
	list.display()
	
	
}
