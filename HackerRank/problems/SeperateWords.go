package main

import (
	"fmt"
	"strings"
)

func main() {

	SeperateWords("Hello   every one out there")

}

func SeperateWords(sentence string) {

	s := strings.Split(sentence, " ")

	for _, v := range s {

		
		fmt.Println(v)
	}

}
