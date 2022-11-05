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

// package main

// import (
// 	"fmt"

// )

// func main() {

// name := "sehtu"

// fmt.Printf("Type of name is %T and the first character is %s",name[4:5],name[4:5])

// }

// func main() {

// 	var input []int

// 	input = append(input, 4)

// 	fmt.Println(len(input))
// }








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
//    var data int
// 	fmt.Println("Enter the data :")
//     fmt.Scan(&data)
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


func (list *SinglyLinkedList) deleteDuplicate() {
    
	
	temp := list.head
	key := list.head
	prev := list.head

	for key != nil  {

		prev = key
		temp = key.Next

		for temp != nil {
			if temp.Data == key.Data {

				prev.Next = temp.Next


	          if temp == list.tail {
				list.tail = prev
			  }
			}

			prev = temp
			temp = temp.Next

		}


		key = key.Next
	}
    
}


func main () {


	s := SinglyLinkedList{}


	s.addNode(1)
	s.addNode(2)

	s.addNode(3)

	s.addNode()
	s.addNode(5)
	s.addNode(5)
	s.display()

	fmt.Print("rfkgndalkghadf;pkgjhdalkgjdnhlgkdja")
	s.deleteDuplicate()
	s.display()
}