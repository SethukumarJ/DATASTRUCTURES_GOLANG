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
		fmt.Print(temp.Data," ")
		temp = temp.Next
	}

}


// ===================================================================functionsssssssf========================================================

func (s *SinglyLinkedList) moreOperation() {
	fmt.Print("\nDuplicate of specific Data --> 'D' \nDelete duplicate of even -->'E' \nDelete duplicate of Odd --> 'O' \nEnter  'A' to delete all the duplicates : " )
	s.duplicateDeletetion()
}
func operations() {
	fmt.Print("\nAddData --> 1 \nDelete --> 2 \nInsert -->3\nDisplay --> 4 \nShow more opetations --> 5 \nEnter 6 to exit :  ")
}




func (list *SinglyLinkedList) insert (data int,search int, counts int) {
    newNode := new(Node)
	newNode.Data = data
	temp := list.head
	
	count := 0

	
	if search == list.head.Data {
		
		newNode.Next = temp 
		list.head = newNode
	    return
	}

	
	for temp.Next != nil {

		if temp.Next.Data== search {

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

func (list *SinglyLinkedList) deleteAllDuplicate() {
	temp := list.head
	key := list.head
	prev := list.head

	var t bool
   
	for key != nil  {

		prev = key
		temp = key.Next
		
		for temp != nil {
			
		   t = temp.Data == key.Data

			if t {

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

 func (list *SinglyLinkedList) deleteDupOf(data int) {

   temp := list.head
   key := list.head
   prev := list.head


 
	for key != nil  {

		prev = key
		temp = key.Next
		
		for temp != nil {
			t := temp.Data == data
			
			if t {
 
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


 func (list *SinglyLinkedList) deleteEven() {
	
    
       
	temp := list.head
	key := list.head
	prev := list.head

	var t bool
   
	for key != nil  {

		prev = key
		temp = key.Next
		
		for temp != nil {
			
		   t = temp.Data%2==0&&key.Data==temp.Data

			if t {

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

 func (list *SinglyLinkedList) deleteOdd() {
	temp := list.head
	key := list.head
	prev := list.head

	var t bool
   
	for key != nil  {

		prev = key
		temp = key.Next
		
		for temp != nil {
			
		   t = temp.Data%2!=0&&key.Data==temp.Data

			if t {

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


 var choice,count,before int

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
			var data int
			fmt.Println("Enter the data to insert :")
			fmt.Scan(&data)
			fmt.Println("Before the data : ")
			fmt.Scan(&before)
			fmt.Println("Enter the count of the data : ")
			fmt.Scan(&count)
			
			s.insert(data,before,count)
		case 4:
		 	s.display()
			break
		case 5:
		 	s.moreOperation()
			break
		case 6:
         break   
	}

	if choice == 6 {
		break
	}
     

 }

}


func main () {
   ListOperations()
}



