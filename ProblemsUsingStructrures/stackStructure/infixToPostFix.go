package main

import "fmt"

type stack struct {
	stackarr []string
	top      int
}

func (s *stack) stack() {

	s.top = -1
}

func (s *stack) push(data string) {

	s.stackarr = append(s.stackarr, data)
	s.top++
}

func (s *stack) pop() string {
	
	s.top--
	return s.stackarr[s.top+1]
}
func (s *stack) isEmpty() bool {
	if s.stackarr == nil {
		return true
	} else {
		return false
	}
}

func (s *stack) peak() string {

	return s.stackarr[s.top]
}

func isLetterOrDigit(input string) bool {

	if input == "+" || input == "-" || input == "*" || input == "/" {

		return false

	} else {

		return true

	}

}

func precedence(character string) int {

	switch character {

	case "+":
	case "-":
		return 1

	case "*":
	case "/":
		return 2

	case "^":
		return 3
	}
	return 0
}

func (s *stack) infixToPostfix(exp string) {
	expression := exp

	var finalExp string
	var character string
	for i := 0; i < len(expression); i++ {
		character = exp[i : i+1]

		if isLetterOrDigit(character) {
			finalExp += expression
		} else {

			precedence := precedence(character) > precedence(s.peak())

			if character == ")" {
				
				for s.peak() != "(" {
					finalExp += s.pop()
				}
				finalExp += s.pop()
				finalExp += ")"

			} else if character == "(" {

				s.push("(")

			} else if s.isEmpty() || character == "(" {
				
				s.push(character)

			}else if !s.isEmpty() && precedence == true {

				s.push(character)

			} else {
				for precedence == false {
					finalExp += s.pop()
				}
				s.push(character)
			}
		}

	}

	for !s.isEmpty() {
		finalExp += s.pop()
	}

	fmt.Println(finalExp)

}


func main() {
	
	convert := stack{}

	convert.infixToPostfix("(a+b^x)*(c+d)")

}
