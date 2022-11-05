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


// ===================================================================functionsssssssf========================================================

func (s *SinglyLinkedList) moreOperation() {
	fmt.Print("\nEnter 'D' to choose a data to delete the duplicate,\nEnter E to delete all the duplicate of even,\nEnter 'O' to delete the duplicate of odd, \nEnter  'A' to delete all the duplicates :" )
	s.duplicateDeletetion()
}
func operations() {
	fmt.Print("\nEnter 1 to addData, \nEnter 2 to delete, \nEnter 3 to display, \nEnter 4 to show more opetations, \nEnter 5 to exit :")
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


func (list *SinglyLinkedList) deleteAllDuplicate() {
    
	
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

 func (s *SinglyLinkedList) deleteDupOf(data int) {
   t :=  subOperations()
	currentNode := s.head

    if t {

        for  currentNode != nil {
			if data == currentNode.Data {
                 s.delete(currentNode.Data)

				 currentNode = currentNode.Next
			}
			currentNode = currentNode.Next
		}   

	} else {

		for currentNode != nil {
			if data == currentNode.Data {
				s.delete (currentNode.Data)
				break
			} else {
				currentNode = currentNode.Next
			}
		}
	}

 }


 func (s *SinglyLinkedList) deleteEven() {
	t :=  subOperations()
    
	currentNode := s.head 

	   if t {
			 for currentNode != nil {

				if currentNode.Data%2 == 0 {
					s.delete(currentNode.Data)
					currentNode = currentNode.Next
				}
				currentNode = currentNode.Next
			 }
	   } else {
			for currentNode != nil {
				if currentNode.Data%2 == 0 {
					s.delete(currentNode.Data)
					break
				} else {
					currentNode = currentNode.Next
				}
			}
	   }
 }

 func (s *SinglyLinkedList) deleteOdd() {

	t :=  subOperations()
	currentNode := s.head 

	   if t {
			 for currentNode != nil {

				if currentNode.Data%2 != 0 {
					s.delete(currentNode.Data)
					currentNode = currentNode.Next
				}
				currentNode = currentNode.Next
			 }
	   } else {
			for currentNode != nil {
				if currentNode.Data%2 != 0 {
					s.delete(currentNode.Data)
					break
				} else {
					currentNode = currentNode.Next
				}
			}
	   }
 }



func (s *SinglyLinkedList) duplicateDeletetion() {
	
  
	var choice string
	var data int
    fmt.Scan(&choice)
    
	switch choice {
	case "D":
		fmt.Println("Enter the data : ")
		fmt.Scan(&data)
		s.deleteDupOf(data)
	case "E":
		s.deleteEven()

	case "O":
		s.deleteOdd()
	case "A":
		s.deleteAllDuplicate()
}


}

func ListOperations() {



 s := SinglyLinkedList{}


 var choice int

for {
	operations()
	fmt.Scan(&choice)

	switch choice {
		
		case 1:
		 	s.addNode()
            break
		case 2:
		 	var data int
			fmt.Print("Enter the data to delete :")
			fmt.Scan(&data)
			s.delete(data)
			break
		case 3:
		 	s.display()
			break
		case 4:
		 	s.moreOperation()
			break
		case 5:
         break   
	}

	if choice == 5 {
		break
	}
     

 }

}


func main () {
   ListOperations()
}



