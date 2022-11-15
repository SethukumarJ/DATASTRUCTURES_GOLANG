package main

import "fmt"

//hash table structure
const size =7
type hashtable struct{
	array [size] *bucket
}




//bucket is a linked list in each slot of the 
type bucket struct {
	head *bucketnode

}
type bucketnode struct{
	key string
	next *bucketnode
}
//insert will take  in a key and add it to the hash table array
 func (h *hashtable) insert(key string){
 	index:=hash(key)                       //the value of hash function remainder
	h.array[index].insert(key)

 }
//search will take in a key and return if that is stored  in the hash table
func (h *hashtable) search(key string)bool {
	index:=hash(key)
   return h.array[index].search(key)
	
}
// delete will take in a key and delete it from the hash table
func (h *hashtable) delete(key string){
	index:=hash(key)
	h.array[index].delete(key)

	
}
// func (h *hashtable) display(){
// 	Init().display()
// }





//insert will take in a key , create a node with the key and insert  the node inthe bucket
func (b *bucket) insert(k string){
	if !b.search(k){
	newnode:=&bucketnode{key:k}
	newnode.next=b.head
	b.head=newnode
	}else{
		fmt.Println(" it is already exits")
	}
}

//search will tAke in a key , create a node with the key and insert the node to the bucket
func (b *bucket )search(k string) bool{
	currentnode :=b.head 
	for currentnode!=nil{
		if currentnode.key==k{
			return true
		} 
		currentnode=currentnode.next
	}
	return false
}
// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(k string){
	previousnode:=b.head

	if b.head.key==k{
		b.head=b.head.next
		return
	}
	for previousnode.next!=nil{
		if previousnode.next.key==k{
			//delete
			previousnode.next=previousnode.next.next
		}
		previousnode=previousnode.next


	}
}

//hash func 
func hash(key string)int {  // the string getting is 
	sum:=0
	for _,v:=range  key {   // the string is splitted into alphabets 
		sum+=int(v)          //store the alphabets integer value 

	}
	return sum %size           //returning the sum value modulus of array size 
	                            //to get  moduls
}

// init will create a bucket in each slot of the hash table
func Init() *hashtable{
result:=&hashtable{}
for i:= range result.array{
	result.array[i]=&bucket{}
}
return result
}
// func (b *bucket) display(){
// 	temp:=b.head
// 	for temp!=nil{
// 		fmt.Println(temp.key)
// 		temp=temp.next
// 	}

// }

func main (){

	 Hashtable:=Init()
	 fmt.Println(Hashtable)
	 list:=[]string{
		"eric",
		//"lee",
		//"kumar",
		//"ramanujan",
		//"alberteinsteen",
		//"prasanth",
		// "kyle",
		// "leonalmessi",
		// "ronaldo",
		// "ram",

	 }
	 //Hashtable.display()
	 for _,v:=range list{
		Hashtable.insert(v)
	 }
 Hashtable.delete("ronldo")
	 fmt.Println(Hashtable.search("lee"))
	 
for i,v :=range Hashtable.array {
	fmt.Println(i,v)

}
	

}