package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	top *Node
}

func (list *Stack) push(data int) {
	newNode := new(Node)
	newNode.Data = data

	if list.top == nil {
		list.top = newNode
	} else {
		newNode.Next = list.top
		list.top = newNode
	}

}

func (list *Stack) pop() {

	if list.top == nil {
		fmt.Println("The stack is empty!")
	} else {
		list.top = list.top.Next
	}

}



func (list *Stack) display() {
   
	currentNode := list.top

	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

}

func main() {

	list := Stack{}

	list.push(1)
	list.push(2)
	list.push(3)
	list.push(4)
	list.push(5)
	list.pop()
	list.display()

}
