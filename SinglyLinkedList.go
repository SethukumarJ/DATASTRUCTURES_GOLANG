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

func (list *SinglyLinkedList) delete(data int){
   temp := list.head

   for {
	if(temp.Next.Data ==data){
		temp.Next = temp.Next.Next
		break
	}
	temp= temp.Next
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

func main() {

	list := SinglyLinkedList{}
	list.addNode(20)
	list.addNode(21)
	list.addNode(22)
	list.addNode(23)
	list.addNode(24)
	list.delete(24)
	list.display()

}
