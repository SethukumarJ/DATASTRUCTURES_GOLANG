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


func (tree *ProgressCard) delete() {

	var name string
	fmt.Print("Enter the name of student to remove : ")
	fmt.Scan(&name)
	tree.deletehelper(tree.root, name)
}


func (tree *ProgressCard) deletehelper(temp *Node,name string) {

	

	for temp != nil {

		if temp.name == name {

			tree.delete(temp.mark)
			return

		} 

			tree.deletehelper(temp.lchild,name)
			tree.deletehelper(temp.rchild,name)
		

	}

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
        temp := key.lchild
        for temp.rchild != nil {
            prev = temp
            temp = temp.rchild
        }
        if prev == key {
            prev.lchild = temp.lchild
        } else {
            prev.rchild = temp.lchild
        }
        return temp.mark
    }

	func getMin(key *Node) int {
        prev := key
        temp := key.rchild
        for temp.lchild != nil {
            prev = temp
            temp = temp.lchild
        }
        if prev == key {
            prev.rchild = temp.rchild
        } else {
            prev.lchild = temp.rchild
        }
        return temp.mark
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