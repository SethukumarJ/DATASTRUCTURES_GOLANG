package main

import "fmt"

type ProgressCard struct {
	root *Node
}

type Node struct {
	mark   int
	name   string
	lchild *Node
	rchild *Node
	parent *Node
}

func (tree *ProgressCard) addNode() {

	var name string
	var mark int
	newNode := new(Node)

	fmt.Print("\nEnter the name of the student to add : ")
	fmt.Scan(&name)
	fmt.Print("\nEnter the mark of ",name," :")
	fmt.Scan(&mark)
	newNode.name = name
	newNode.mark = mark
	if tree.root == nil {

		tree.root = newNode
		return

	}

	temp := tree.root

	for temp != nil {

		if newNode.mark >= temp.mark {

			if temp.lchild == nil {

				temp.lchild = newNode
				newNode.parent = temp
				return 
			}

			temp = temp.lchild

		} else if newNode.mark < temp.mark {

			if temp.rchild == nil {

				temp.rchild = newNode
				newNode.parent = temp
				return

			}

			temp = temp.rchild

		}

	}

}


func (tree *ProgressCard) inOrder () {

	fmt.Println("=======RANK LIST========")
	tree.inOrderHelper(tree.root)
	
}


func (tree *ProgressCard) inOrderHelper (temp *Node) {

	if temp != nil  {

		tree.inOrderHelper(temp.lchild)
		fmt.Println(temp.name, " - ",temp.mark)
		tree.inOrderHelper(temp.rchild)
	}

}



func main () {

	count :=1
	tree := ProgressCard{}

	for count == 1 {
            
		tree.addNode()

		fmt.Println("Enter 1 to continue")
		fmt.Scan(&count)
    
	}

	tree.inOrder()
	return
}