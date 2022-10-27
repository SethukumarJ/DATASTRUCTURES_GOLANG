package main

import "fmt"

type Queue struct {
	front *Node
	rear  *Node
}

type Node struct {
	Data int
	Next *Node
}

func (que *Queue) append(data int) {
	newNode := new(Node)

	newNode.Data = data
	newNode.Next = nil

	if que.front == nil {
		que.front = newNode
	} else {
		que.rear.Next = newNode
	}

	que.rear = newNode

}

func (que *Queue) eliminate() {
	if que.front == nil {

		fmt.Println("Queue is empty")

	} else {
		que.front = que.front.Next
	}
}


func (que *Queue) display() {
    
	if que.front ==nil {
		fmt.Println("The que is empty!")
	}else {
		currentNode := que.front
    
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

	}
	
}

func main() {

	q := Queue{}
	q.append(1)
	q.append(2)
	q.append(3)
	q.append(4)
	q.eliminate()
	q.display()

}
