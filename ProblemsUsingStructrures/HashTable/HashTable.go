package main

import "fmt"

type Node struct {
	Key   string
	Value interface{}
	Next  *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

func (list *SinglyLinkedList) addNode(key string, value interface{}) {
	newNode := new(Node)
	newNode.Value = value
	newNode.Key = key
	newNode.Next = nil

	if list.head == nil {
		list.head = newNode

	} else {
		list.tail.Next = newNode
	}

	list.tail = newNode

}
func (list *SinglyLinkedList) display(index int) {
	temp := new(Node)
	if(list.head ==nil) {
		fmt.Println("Empty list!")
	}

	temp = list.head

	for temp !=nil {
		fmt.Println("[",index,"]",temp.Key,":",temp.Value)
		temp = temp.Next
	}

}

const tableSize = 10

type HashTable struct {
	Pair [tableSize]*SinglyLinkedList
}

func (ht *HashTable) hash(key string) int {
	sum := 0
	for i := 0; i < len(key); i++ {

		sum += int(key[i])
	}
	return sum % tableSize
}

func (ht *HashTable) Insert(key string, value interface{}) {

	index := ht.hash(key)

	if ht.Pair[index] == nil {
		list := new(SinglyLinkedList)
		ht.Pair[index] = list
		list.addNode(key, value)
	} else {
		ht.Pair[index].addNode(key, value)
	}
}

func (ht *HashTable) Search(key string) interface{} {

	index := ht.hash(key)

	if ht.Pair[index] == nil {
		return "Searched key is not present!"
	} else {
		temp := ht.Pair[index].head
		for temp != nil {
			if temp.Key == key {
				return temp.Value
			}
			temp = temp.Next
		}
	}

	return "No value"
}


func (ht *HashTable) displayHashTable () {
	var i =0
	for i<tableSize{

		if ht.Pair[i] != nil {
			ht.Pair[i].display(i)
		}	
		i++
	}
}

func main() {

	ht := HashTable{}

	ht.Insert("sethu", 12)
	ht.Insert("rahul", 13)
	ht.Insert("gokul", 14)
	ht.Insert("ramu", 15)
	ht.Insert("gopu", 16)
	ht.Insert("gopv",17)

	ht.displayHashTable()

	searchWord := "akdfksd"
	search := ht.Search(searchWord)

	fmt.Println(searchWord , " : ", search)
}
