package main

import "fmt"

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoublyLinkedList) addNode(data int) {

	newNode := new(Node)
	newNode.Data = data
	newNode.Next = nil

	if list.head == nil {
		list.head = newNode
	} else {
		
		list.tail.Next = newNode
		newNode.Prev = list.tail
	}

	list.tail = newNode
}

func (list *DoublyLinkedList) displayForward() {
	currentNode := new(Node)

	currentNode = list.head

	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

}

func (list *DoublyLinkedList) displayBackward() {
	currentNode := new(Node)

	currentNode = list.tail

	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Prev
	}
}

func (list *DoublyLinkedList) insertBefore(Data int, Before int) {
	temp := new(Node)
	newNode := new(Node)
	newNode.Data = Data
    
	temp = list.head
	if list.head.Data == Before {
		list.head.Prev = newNode
		newNode.Next = list.head
		list.head = newNode
	} else {
		for temp.Data != Before {

			if temp.Next.Data == Before {
				newNode.Next = temp.Next
				newNode.Prev = temp
				temp.Next = newNode
				break
			}
			temp = temp.Next
		}
	}

}

func (list *DoublyLinkedList) delete(data int){

	currentNode := new(Node)
    
	currentNode = list.head

	for currentNode != list.tail {

		if(currentNode.Data == data) {

			currentNode.Prev.Next = currentNode.Next
			currentNode.Next.Prev = currentNode.Prev
			
		    break
			
		}
		currentNode = currentNode.Next
    }

}

func main() {

	list := DoublyLinkedList{}
	list.addNode(4)
	list.addNode(3)
	list.addNode(9)
	list.addNode(0)
	list.insertBefore(6,0)
	list.delete(6)
	list.displayForward()
}
