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

			tree.remove(temp.mark)
			return

		} 

			tree.deletehelper(temp.lchild,name)
			tree.deletehelper(temp.rchild,name)
		

	}

}

func (tree *ProgressCard) remove (mark int) {

	
	temp := tree.root
	if temp == nil {

		fmt.Println("Tree is empty !...")
		return

	}

	for temp != nil {

		if mark > temp.mark {

			temp = temp.lchild

		} else if mark < temp.mark {

			temp = temp.rchild

		} else {

				if temp.rchild!=nil {

                	temp.mark = getMin(temp)

				} else if temp.lchild!=nil {

                	temp.mark = getMax(temp)

				} else {

                	temp.mark = getMin(temp)
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

func (tree *ProgressCard) display () {

	fmt.Println("=======RANK LIST========")
	tree.displayHelper(tree.root)
	
}


func (tree *ProgressCard) displayHelper (temp *Node) {

	if temp != nil  {

		tree.displayHelper(temp.lchild)
		fmt.Println(temp.name, " - ",temp.mark)
		tree.displayHelper(temp.rchild)
	}

}






func main () {

	count :=1
	
	tree := ProgressCard{}

	for count != 4 {
		fmt.Println("\nadd --> 1 \nremove --> 2 \ndisplay --> 3 \nexit -->4")
		fmt.Scan(&count)

		switch count {

		case 1:
			tree.addNode()
		case 2:
			tree.delete()
		case 3:
			tree.display()
		case 4:
			count = 4
		}
            
		
	}

}