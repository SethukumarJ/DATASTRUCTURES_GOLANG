package sethu

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	top *Node
}

func (list *Stack) Push(data int) {
	newNode := new(Node)
	newNode.Data = data

	if list.top == nil {
		list.top = newNode
	} else {
		newNode.Next = list.top
		list.top = newNode
	}

}

func (list *Stack) Pop() {

	if list.top == nil {
		fmt.Println("The stack is empty!")
	} else {
		list.top = list.top.Next
	}

}

func (list *Stack) Display() {

	currentNode := list.top

	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

}


