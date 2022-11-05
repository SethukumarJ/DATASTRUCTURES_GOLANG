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

func (list *SinglyLinkedList) addNode() {
   var data int
	fmt.Println("Enter the data :")
    fmt.Scan(&data)
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

	
   temp := list.head

   for {
	if(temp.Next.Data ==data){
		temp.Next = temp.Next.Next
		break
	}
	temp= temp.Next
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

func moreOperation() {
	fmt.Print("\nEnter 'D' to choose a data to delete the duplicate,\nEnter E to delete all the duplicate of even,\nEnter 'O' to delete the duplicate of odd, \nEnter  'A' to delete all the duplicates" )
}
func operations() {
	fmt.Print("\nEnter 1 to addData, \nEnter 2 to delete, \nEnter 3 to display, \nEnter 4 to show more opetations, \nEnter 5 to exit")
}

func subOperations() bool {
	var option string
	fmt.Print("Enter 'F' to delete the first duplicate, \nEnter 'A' to delete all duplicates : ")
    fmt.Scan(&option)
	if option == "F" {
		return false
		
	} else if option == "A" {
		return true
	} else {
		panic("Give correct letter given in the option :")
	}
	
}


 func (s *SinglyLinkedList) deleteAllDuplicate() {

    currentNode := s.head

	for currentNode != s.tail {
           successorNode := currentNode.Next
		for successorNode != nil {

			if currentNode.Data == successorNode.Data {
				s.delete(successorNode.Data)
				successorNode = successorNode.Next
			}
		}
		currentNode = currentNode.Next
	}

 }   

 func (s *SinglyLinkedList) deleteSpecificData() {

	

 }



func duplicateDeletetion() {
   s := SinglyLinkedList{}
	var choice,choice2 string
	var data int
    operations()
	fmt.Scan(&choice)
    
	switch choice {
	case "D":
		fmt.Println("Enter the data : ")
		fmt.Scan(&data)
	case "E":
		subOperations() 
		fmt.Scan(&choice2)
	case "O":
		subOperations() 
		fmt.Scan(&choice2)
	case "A":
		s.deleteAllDuplicate()

}


}

func deleteDuplicate(data int) {

 s := SinglyLinkedList{}

 var choice1,choice2,choice3 int



 for choice1 !=5 {
       
	operations()
	fmt.Scan(&choice1)

	switch choice1 {
	case 1:
		s.addNode()

	case 2:
		var data int
		fmt.Print("Enter the data to delete :")
		fmt.Scan(&data)
		s.delete(data)

	case 3:
		s.display()
	case 4:
		moreOperation()
	case 5:
		break
	}
     

 }



}