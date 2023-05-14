package main

import (
	"fmt"
)

func AlphaNumeric(word string) {

	for i:=0; i<len(word); i++ {

		if (word[i] >= 65 && word[i] <= 122) || (uint8(word[i]) >= 49 && uint8(word[i]) <= 57) {

			fmt.Println(string(word[i]))
		}

	}

}

func main() {

	AlphaNumeric("ABZ1345*#.!a09")

	// s := "l"

	// fmt.Println(s[0])

	
}
