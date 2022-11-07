package main

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

func (list *SinglyLinkedList) addNode(data string) {
	newNode := new(Node)
	newNode.Data = data
	newNode.Next = nil

	if list.head == nil {
		list.head = newNode

	} else {
		list.tail.Next = newNode
	}

	list.tail = newNode

}
func (list *SinglyLinkedList) delete(data string) {

	if list.head != nil && list.head.Data == data {
		list.head = list.head.Next
		return
	}
	temp := list.head
	for temp != nil {
		if temp.Next.Data == data {
			if temp.Next == list.tail {
				list.tail = temp
				list.tail.Next = nil
			} else {
				temp.Next = temp.Next.Next

				return
			}
		}

		temp = temp.Next
	}

}

func (list *SinglyLinkedList) display() {
	temp := new(Node)
	if list.head == nil {
		fmt.Println("Empty list!")
	}

	temp = list.head

	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}

}

func (list *SinglyLinkedList) insert(data int, search string, counts int) {
	newNode := new(Node)
	newNode.Data = search
	temp := list.head

	count := 0

	if search == list.head.Data {

		newNode.Next = temp
		list.head = newNode
		return
	}

	for temp.Next != nil {

		if temp.Next.Data == search {

			count++

			if count == counts {
				newNode.Next = temp.Next
				temp.Next = newNode
				return
			}

		}
		temp = temp.Next

	}
}

func (list *SinglyLinkedList) paliondromeCheck() {

	var data string
	fmt.Print("Enter the string : ")
	fmt.Scan(&data)
	length := len(data)
	var name string
	count := 1
	
	
	
	for i := 0; i < length; i++ {

		list.addNode(string(data[i]))

	}

	temp := list.head

	for count <= length{

		if count == length {

			name += temp.Data
			
			temp = list.head
			count = 1
			length--
		} else {

		temp = temp.Next
		 count++

		}
		
		
		
	}

   if name == data {
	fmt.Println("paliondrome!")
   } else {
	fmt.Println("not paliondrome!")
   }

}

func main() {

	list := SinglyLinkedList{}

	list.paliondromeCheck()

}
