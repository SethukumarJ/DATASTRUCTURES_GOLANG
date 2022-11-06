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




func getMax(temp *Node) int {

	q := temp
	temp = temp.left

	for temp != nil {

		q = temp
		temp = temp.right

	} 
	
	if temp.left != nil {

		q.right = temp.left

	}

	return temp.Data

}

func getMin(temp *Node) int {

	q := temp
	temp = temp.right

	for temp != nil {

		q = temp
		temp = temp.left

	} 
	
	if temp.left != nil {

		q.left = temp.right

	} 
	return temp.Data

}


func (tree *BinarySearchTree) inOrder () {

	fmt.Print("Inorder : ")
	tree.inOrderHelper(tree.root)
	
}


func (tree *BinarySearchTree) inOrderHelper (temp *Node) {

	if temp != nil  {

		tree.inOrderHelper(temp.left)
		fmt.Print(temp.Data," ")
		tree.inOrderHelper(temp.right)
	}

}


func (tree *BinarySearchTree) preOrder () {

	fmt.Print("PreOrder : ")
	tree.preOrderHelper(tree.root)
	
}


func (tree *BinarySearchTree) preOrderHelper (temp *Node) {

	if temp != nil  {

		fmt.Print(temp.Data," ")
		tree.preOrderHelper(temp.left)
		tree.preOrderHelper(temp.right)
	}

}


func (tree *BinarySearchTree) postOrder () {

	fmt.Print("PostOrder : ")
	tree.postOrderHelper(tree.root)
	
}


func (tree *BinarySearchTree) postOrderHelper (temp *Node) {

	if temp != nil  {

		
		tree.postOrderHelper(temp.left)
		tree.postOrderHelper(temp.right)
		fmt.Print(temp.Data," ")

	}

}



func main () {

    tree := BinarySearchTree{}

	tree.addNode()
	tree.addNode()
	tree.addNode()
	tree.addNode()
	tree.addNode()
	tree.addNode()
	tree.inOrder()
	tree.postOrder()
	tree.preOrder()

}


