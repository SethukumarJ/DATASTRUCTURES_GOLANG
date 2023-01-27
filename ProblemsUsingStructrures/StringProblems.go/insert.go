package main

import "fmt"

func main() {
 insert(1 , "hellz")
}

func insert(position uint8, word string) {

	var newWord string

	for i := 0; i < len(word); i++ {

		if word[i] + position > 122 {
			newWord += string(96+(position-(122-word[i])))
		} else {
			newWord += string(word[i] + position)
		}
		
	}

	fmt.Println(newWord)
}
