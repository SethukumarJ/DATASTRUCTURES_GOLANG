package main

import "fmt"

func main() {

	var s = [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	

	element := 4


	for i:=0; i< len(s); i++{

      if s[i] == element {
        
	for j:= i; j< len(s)-1; j++ {
		 
		s[j] = s[j+1]

	}

	  }

	}
	fmt.Println(s)
}
