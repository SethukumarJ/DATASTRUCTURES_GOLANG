package main

import (
	"fmt"
)


type Node struct {

	Data int
	Next *Node
	prev *Node
}

type PriorQueue struct {
    
	front *Node
	rear *Node

}


func (node *PriorQueue) append(data int) {
    
	newNode := new(Node)
	newNode.Data = data
    temp := node.front
	if node.front ==nil {

		node.front = newNode
		node.rear = newNode
    } else if data < node.front.Data {

		newNode.Next = node.front
		node.front.prev =newNode
		node.front = newNode

	} else if data > node.rear.Data {
        
		newNode.prev = node.rear
        node.rear.Next = newNode
        node.rear = newNode

    } else {
		for temp != nil  {
           if temp.Data > data {
                newNode.Next = temp
				newNode.prev = temp.prev
				temp.prev.Next = newNode
				temp.prev =newNode
				break
		   }
		   temp = temp.Next
		}
	}

	// node.rear = newNode
}

func (node *PriorQueue) display() {
    
	if node.front ==nil {
		fmt.Println("The que is empty!")
	}else {
		currentNode := node.front
    
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

	}
	
}


func main() {
	q := PriorQueue{}
	q.append(1)

	q.append(5)
	q.append(2)
	q.append(4)
	q.append(3)
	
	q.display()
	

}