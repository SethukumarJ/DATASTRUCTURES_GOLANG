//Problem using linear search.
// Find the index of specific character in slice

package main

import "fmt"

func main() {

	index := FindIndex("sethu", "t")

	fmt.Println(index)

}


func FindIndex(input string, character string) int {
      
	for i,v := range input {
		if string(v) == character {
			return i
		}
	}
	return -1
}