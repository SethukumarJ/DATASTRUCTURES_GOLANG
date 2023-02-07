package main

import (
	"fmt"
	"math"
)

type Node struct {
	Data   int
	right  *Node
	left   *Node
	parent *Node
}

type BinarySearchTree struct {
	root *Node
}

func (tree *BinarySearchTree) addNode(data int){

	newNode := new(Node)
	newNode.Data = data
	
	if tree.root == nil {
		tree.root = newNode 
		newNode.parent, newNode.left, newNode.right = nil, nil ,nil
		return
	}

	temp := tree.root
	for temp!= nil {

		if data < temp.Data {

			if temp.left == nil {
				temp.left = newNode
				newNode.parent = temp
				return
			}
			temp = temp.left
		} else {

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

	for temp != nil {

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

func (tree *BinarySearchTree) delete(data int) {
	temp := tree.root
	prev := temp

	if temp == nil {
		fmt.Println("Tree is empty!")
		return
	}

	for temp != nil {

		if data < temp.Data {
			prev = temp
			temp = temp.left
		} else if data > temp.Data {
			prev = temp
			temp = temp.right
		} else {
			if temp.right != nil {
				temp.Data = getMin(temp)
			} else if temp.left != nil {
				temp.Data = getMax(temp)
			} else {
				if temp.Data < prev.Data {
					prev.left = nil
				} else {
					prev.right = nil
				}
			}

			fmt.Println("deleted!")
			return
		}

		fmt.Println("data not found!")
	}
}

// func replace(temp *Node, prev *Node) {

// }

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

func (tree *BinarySearchTree) sumOfTree() {

	sum := tree.sumHelper(tree.root)
	fmt.Println("The sum is : ", sum)

}

var sum int

func (tree *BinarySearchTree) sumHelper(temp *Node) int {

	if temp != nil {

		tree.sumHelper(temp.left)
		sum += temp.Data
		tree.sumHelper(temp.right)

	}

	return sum

}

func (tree *BinarySearchTree) FindNearest(data int) {

	near := tree.FindNearestHelper(tree.root, data)
	fmt.Println("The nearest data is : ", near)

}

var near int
var diff int

func (tree *BinarySearchTree) FindNearestHelper(temp *Node, data int) int {

	if temp == tree.root {
		diff = int(math.Abs(float64(temp.Data - data)))

		near = temp.Data
	}
	if temp != nil {

		tree.FindNearestHelper(temp.left, data)
		if int(math.Abs(float64(temp.Data-data))) <= diff {
			diff = int(math.Abs(float64(temp.Data - data)))
			near = temp.Data
		}
		tree.FindNearestHelper(temp.right, data)

	}

	return near

}

func (tree *BinarySearchTree) inOrder() {

	fmt.Print("Inorder : ")
	tree.inOrderHelper(tree.root)

}

func (tree *BinarySearchTree) inOrderHelper(temp *Node) {

	if temp != nil {

		tree.inOrderHelper(temp.left)
		fmt.Print(temp.Data, " ")
		tree.inOrderHelper(temp.right)
	}

}

func (tree *BinarySearchTree) preOrder() {

	fmt.Print("PreOrder : ")
	tree.preOrderHelper(tree.root)

}

func (tree *BinarySearchTree) preOrderHelper(temp *Node) {

	if temp != nil {

		fmt.Print(temp.Data, " ")
		tree.preOrderHelper(temp.left)
		tree.preOrderHelper(temp.right)
	}

}

func (tree *BinarySearchTree) postOrder() {

	fmt.Print("PostOrder : ")
	tree.postOrderHelper(tree.root)

}

func (tree *BinarySearchTree) postOrderHelper(temp *Node) {

	if temp != nil {

		tree.postOrderHelper(temp.left)
		tree.postOrderHelper(temp.right)
		fmt.Print(temp.Data, " ")

	}

}

func main() {

	tree := BinarySearchTree{}
	tree.addNode(3)
	tree.addNode(3)
	tree.addNode(3)
	tree.addNode(3)
	tree.inOrder()

}




