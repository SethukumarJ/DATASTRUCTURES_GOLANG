package main

import "fmt"

type Node struct {
	Data rune
	Next *Node
}

type Stack struct {
	top *Node
}

func (list *Stack) push(data rune) {
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

func (list *Stack) peak() rune {
	
	if list.top == nil {
		panic("This stack is empty")
		
	} 
		return list.top.Data
	
}



func (list *Stack) display() {
   
	currentNode := list.top

	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

}


func PaliodromeCheck(data string) {
    
	list := Stack{}
    var count int
	for _, pdata := range data {
		list.push(pdata)
	}

	for _, pop := range data {
     
		var peak rune
        peak = list.peak()
		if peak != pop {
			count++
			break
		} else {
			list.pop()
		}
	}

	if count != 0{
		fmt.Println("Not a paliondrome")
	} else {
		fmt.Println("This string is  a paliondrome")
	}

	

}


func main () {

	var data string
    
	fmt.Print("Enter the string : ")
	fmt.Scan(&data)
    PaliodromeCheck(data)
}

