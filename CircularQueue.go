package main

import "fmt"

type CircularQueue struct {
	front *Node
	rear  *Node
}

type Node struct {
	Data int
	Next *Node
}

func (c *CircularQueue) enQue(data int) {

	newNode := new(Node)
	newNode.Data = data

	if c.front == nil {
		c.front = newNode

	} else {
     
		c.rear.Next = newNode
		newNode.Next = c.front
	   
	}
	
	c.rear = newNode

}

func (c *CircularQueue) deQue() {

	c.front = c.front.Next
    c.rear.Next = c.front
}

func (c *CircularQueue) display() {
	currentNode := c.front
	if c.front ==nil{
		fmt.Println("The que is empty")
	} else {

		for{
			fmt.Println(currentNode.Data)
	
			if currentNode.Next == c.front{
				break
			}
			currentNode = currentNode.Next
		}
		
	}

}

func main() {

	q := CircularQueue{}

	q.enQue(1)
	q.enQue(2)
	q.enQue(3)
	q.enQue(4)
	q.enQue(5)
	q.deQue()
	q.deQue()
	q.display()

}
