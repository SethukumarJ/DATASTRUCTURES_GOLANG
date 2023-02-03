package main

import "fmt"

type stack struct {
	size int
	data []int
	top  int
}


func (s *stack) push(data int) {

	if s.size == s.top-1 {
		fmt.Println("stack overflow")
	}

	if s.top == -1 {
		s.data = append(s.data, data)
		s.top = 0
	} else {

		s.data = append(s.data, data)
		s.top++
	}

}

func (s *stack) pop() {

	if s.top == -1 {
		fmt.Println("stack is empty!")
	} else {
		s.top--
	}
}

func (s *stack) display() {

	for i := 0; i <= s.top; i++ {

		fmt.Println(s.data[i])
	}

}

func main() {

	s := stack{}
	s.top=-1
	s.size = 7
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	s.pop()
	s.display()
}
