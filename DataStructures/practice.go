// package main

// import "fmt"

// type Node struct {
// 	Data int
// 	Next *Node
// }

// type SinglyLinkedList struct {
// 	head *Node
// 	tail *Node
// }

// func (list *SinglyLinkedList) addNode(data int) {

// 	newNode := new(Node)

// 	newNode.Data = data
// 	newNode.Next = nil

// 	if list.head == nil {
// 		list.head = newNode
// 	} else {
// 		list.tail.Next = newNode
// 	}

// 	list.tail = newNode

// }

// func (list *SinglyLinkedList) display() {

// 	currentNode := new(Node)

// 	if list.head == nil {
// 		fmt.Println("List is empty")
// 	}

// 	currentNode = list.head
// 	  for currentNode != nil {
// 		fmt.Println(currentNode.Data)
// 		currentNode = currentNode.Next
// }
// 	  }

package main

import (
	"fmt"

)

// func main() {

// name := "sehtu"

// fmt.Printf("Type of name is %T and the first character is %s",name[4:5],name[4:5])

// }

func main() {

	var input []int

	input = append(input, 4)

	fmt.Println(len(input))
}
