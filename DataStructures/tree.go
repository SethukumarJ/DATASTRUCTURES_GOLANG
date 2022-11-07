public class Tree {

    static class Node {
        int data;
        Node left, right;

        Node(int data) {
            this.data = data;
        }
    }

    Node root;

    public void insert(int data) {

        Node currentNode = root;

        if(currentNode == null){
            root = new Node(data);
            return;
        }

        while (true) {
            if (data < currentNode.data) {

                if (currentNode.left == null) {
                    currentNode.left = new Node(data);
                    break;
                } else {
                    currentNode = currentNode.left;
                }
            } else {
                if (currentNode.left == null) {
                    currentNode.left = new Node(data);
                    break;
                } else {
                    currentNode = currentNode.left;
                }
            }
            
        }
        

    }

    public boolean contain(int data) {

        Node currentNode = root;

        while (currentNode != null) {
            if (data < currentNode.data) {
                currentNode = currentNode.left;
            } else if (data > currentNode.data) {
                currentNode = currentNode.right;
            } else {

                return true;
            }
        }

        return false;

    }

    public void remove(int data) {
        removeHelper(data, root, null);

    }

    public void removeHelper(int data, Node currentNode, Node parentNode){
        currentNode = root;

        while(currentNode != null){
            if(data< currentNode.data){
                
                    parentNode = currentNode;
                    currentNode = currentNode.left;
                }
                
                else if(data>currentNode.data){

                    parentNode = currentNode;
                    currentNode = currentNode.right;
                }

                else{
                    if(currentNode.left != null && currentNode.right != null){
                        currentNode = minValue(currentNode.right);
                        removeHelper(currentNode.data, currentNode.right, currentNode);

                        
                    }
                    else{

                       if(parentNode == null){
                        if(currentNode.right == null){
                            root = currentNode.left;
                        }
                        else{
                            root = currentNode.right;
                        }
                       }
                       else{
                        if(currentNode.right == null){
                            if(currentNode.left != null){
                                parentNode.left = currentNode.left;
                            }
                            else{
                                parentNode.right = currentNode.right;
                            }
                        }
                        else{
                            if(currentNode.right != null){
                                parentNode.right = currentNode.right;
                            }
                            else{
                                parentNode.left = currentNode.left;
                            }

                        }
                       }

                    }
                    break;
                }
                
            }
        }


        
    

    public Node minValue(Node node) {

        node = root;

        if(node.left == null){

            return node;
           
        }
       
else{ 
   return minValue(node.left);
}

    }

    public static void main(String[] args) {
        Tree tree = new Tree();

        tree.insert(10);
        tree.insert(8);
        tree.insert(11);
        tree.remove(8);
       
        
        System.out.println(tree.contain(8));

    }

}
