package main

import "fmt"

func main() {

	reversed := reverseString("hello")
	fmt.Println(reversed)

}

func reverseString(s string) string {
	var reversed string
	length := len(s)

	for i := length-1; i>=0; i-- {
		reversed += string(s[i])
	}

	return reversed
}
