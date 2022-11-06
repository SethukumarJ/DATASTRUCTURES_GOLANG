package main


import (
	"fmt"
)


type Node struct {

	Data int
	right *Node
	left *Node
	parent *Node

}


type BinarySearchTree struct {
	
	root *Node

}


func (tree *BinarySearchTree) addNode () {

	var data int
	newNode := new(Node)

	fmt.Print("\nEnter the data to add :")
	fmt.Scan(&data)

	newNode.Data = data

	if (tree.root == nil) {
		
		tree.root = newNode
		return

	}

	temp := tree.root

	for temp != nil {

		if newNode.Data < temp.Data {

			if temp.left == nil {

				temp.left = newNode
				newNode.parent = temp
				return
			}

			temp = temp.left

		} else if newNode.Data > temp.Data {

			if temp.right == nil {

				temp.right = newNode
				newNode.parent = temp
				return

			}

			temp = temp.right

		}

	}

}


func (tree *BinarySearchTree) contain() {

	var data int
	temp := tree.root

	fmt.Print("\nEnter the data to check whether the tree contains or not : ")
	fmt.Scan(&data)

	if temp == nil {

		fmt.Println("Tree is empty!...")
		return
	}

	for temp != nil  {

		if data < temp.Data {

			temp = temp.left

		} else if data > temp.Data {

			temp = temp.right

		} else {

			fmt.Println("Match found!")
			return

		}
	}

	fmt.Println("No match found!")
	
}


