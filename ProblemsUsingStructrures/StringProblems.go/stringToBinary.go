package main

import "fmt"

func main() {

	word := "sethu"

	for i := 0; i < len(word); i++ {

		fmt.Printf("%08s ", string(byte(word[i])))
	}
}
