
package main 





type BinarySearchTree struct {

	root *Node
}

type Node struct {

	data int
	parent *Node
	left *Node
	right *Node
}



func (tree *BinarySearchTree) insert(data int) {

	newNode := new(Node)
	newNode.data = data
	temp := tree.root

	if temp == nil {
		tree.root = newNode
		return
	}


	for temp != nil {


		if data < temp.data {
			if temp.left == nil
			temp = temp.left 
		}
	}
}




func main(){

}