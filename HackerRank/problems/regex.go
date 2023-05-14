package main

import "fmt"

func main() {

	regex()


}

func regex() {

	var valid bool
	valid = true
	var numberString string
	fmt.Println("Enter the Phone number : ")
	fmt.Scan(&numberString)

	if len(numberString) >10 {

		valid = false
	}

	

	for i, _ := range numberString {

		if !((numberString[i]) >=48 && numberString[i] <= 57) {

			valid = false

		}
	}

	if valid == true {
		fmt.Println("valid")
	} else {
		fmt.Println("not valid")
	}


}
