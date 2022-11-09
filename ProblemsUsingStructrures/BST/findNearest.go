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


func (tree *BinarySearchTree) remove () {

	var data int
	temp := tree.root

	fmt.Print("\nEnter the data to remove : ")
	fmt.Scan(&data)

	if temp == nil {

		fmt.Println("Tree is empty !...")
		return

	}

	for temp != nil {

		if data < temp.Data {

			temp = temp.left

		} else if data > temp.Data {

			temp = temp.right

		} else {

				if temp.right!=nil {

                	temp.Data = getMin(temp)

				} else if temp.left!=nil {

                	temp.Data = getMax(temp)

				} else {

                	temp.Data = getMin(temp)
                	fmt.Println("Data deleted succesfully")
                	return

				}
			}
		}

	}

	func getMax(key *Node) int {
        prev := key
        temp := key.left
        for temp.right != nil {
            prev = temp
            temp = temp.right
        }
        if prev == key {
            prev.left = temp.left
        } else {
            prev.right = temp.left
        }
        return temp.Data
    }

	func getMin(key *Node) int {
        prev := key
        temp := key.right
        for temp.left != nil {
            prev = temp
            temp = temp.left
        }
        if prev == key {
            prev.right = temp.right
        } else {
            prev.left = temp.right
        }
        return temp.Data
    }





	func (tree *BinarySearchTree) findNearest() {

		var data,minValue int
		minValue = 1000
		fmt.Print("Enter the data to search for the nearest : ") 
		fmt.Scan(&data)
		var nearest int
		tree.findNearestHelper(tree.root,data,minValue,&nearest)
	}


	func (tree *BinarySearchTree) findNearestHelper(temp *Node,data int,minValue int,nearest *int) int{

		
		if temp != nil {

			tree.findNearestHelper(temp.left,data,minValue,nearest)

			diff := temp.Data - data 
			if diff < 0 {

				diff = diff * -1
			}

			if diff < minValue {

				minValue = diff
				*nearest = temp.Data
			
			}

			tree.findNearestHelper(temp.right,data,minValue,nearest)

		}

		return *nearest
		

	}