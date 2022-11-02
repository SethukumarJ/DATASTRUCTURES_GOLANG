package main

import "fmt"

type stack struct {
	limit    int
	stackarr []int
	top      int
}

func (s *stack) stack(limit int) {


	s.limit = limit
	s.top = -1
}

func (s *stack) push(data int) {

	s.stackarr = append(s.stackarr, data)
	s.top++
}

func (s *stack) pop() {

	s.stackarr[s.top] = 0
	s.top--

}
func (s *stack) isEmpty() bool {
	if s.stackarr == nil {
		return true
	} else {
		return false
	}
}
func (s *stack) isFull() bool {
	if s.top == s.limit-1 {
		return true
	} else {
		return false
	}
}

func (s *stack) peak() int {

	return s.stackarr[s.top]
}

func main() {
    
	st := stack{}
    limit,choice,data := 0,0,0
	fmt.Println("Enter the limit of the stack : ")
	fmt.Scan(&limit)
    st.stack(limit)
	

	for choice !=4 {
	
     fmt.Println("Enter 1 to push, 2 to pop, 3 to get the peak value and 4 to exit !")
     fmt.Scan(&choice)
	 switch (choice) {
	 case 1:
		fmt.Print("Enter the data : ")
			fmt.Scan(&data)
			st.push(data)
			break
	 case 2:
		if st.isEmpty() == false {
			st.pop()
			break
		} else {
			fmt.Println("stack is empty!")
			break
		}
	case 3:
		data = st.peak()
        fmt.Println(data)
		break

	case 4:
		break
   }

	}

}
