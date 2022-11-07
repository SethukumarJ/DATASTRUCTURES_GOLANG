package main

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type Stack struct {
	top *Node
}

func (list *Stack) push(data string) {
	newNode := new(Node)
	newNode.Data = data

	if list.top == nil {
		list.top = newNode
	} else {
		newNode.Next = list.top
		list.top = newNode
	}

}

func (list *Stack) pop() string {

	temp := list.top

	if list.top == nil {
		fmt.Println("The stack is empty!")
		return ""
	} else {

		list.top = list.top.Next

	}
	return temp.Data

}

func (list *Stack) display() {

	currentNode := list.top

	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

}

func longestValidParentheses(s string) int {

	count := 0
	countdup := 0
	stack := Stack{}

	for i := 0; i < len(s); i++ {
		element := string(s[i])
		stack.push(element)
	}

	temp := stack.top

	for temp != nil {

		character := stack.pop()

		if character == ")" {
			if stack.top.Data == "(" {

				stack.pop()
				count++

				if count > countdup {
					countdup = count
				}

			} else if stack.top.Data == ")" {

				count = 0

			}

		}

		temp = stack.top

	}

	return countdup

}

func main() {

	s := longestValidParentheses("(()()()))")

	fmt.Println(s)
}
