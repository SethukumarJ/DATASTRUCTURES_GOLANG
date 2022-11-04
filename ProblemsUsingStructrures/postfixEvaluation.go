package main

import (
	"reflect"
	"strconv"
	"time"
)

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

func Digit(exp string) bool {
	digitType := 4
	digit, err := strconv.Atoi(exp)
	if err != nil {
		time.Sleep(time.Microsecond)
	}
	if reflect.TypeOf(digit) == reflect.TypeOf(digitType) {
		return true
	} else {
		return false
	}
}
func operation(operator string, a int, b int) int{
	switch operator{
		case "+":
		   return a+b
		
	    case "-":
			return a-b

		case "*":
			return a*b

		case "/":
			return a/b

	}
	return 0;
}


func postfixEvaluator(exp string) {



}