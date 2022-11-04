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

func (s *stack) pop() {
	s.top--
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

func isLetterOrDigin(input string) bool {

	 if input == "+" || input == "-" || input == "*" || input == "/" {

		return true

	 } else {
		
		return false

	 }

}




func (s *stack) infixToPostfix(input string) {
   
       

}


