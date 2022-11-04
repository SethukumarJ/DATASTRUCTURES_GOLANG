package main


import (
	 "fmt"
	)

func Sorting (input []int) []int{
      
  for i := 1; i < len(input); i++ {
	current := i

	for j := i-1; j >=0; j-- {

		if input[j]>input[current] {

			temp := input[j]
			input[j] = input[current]
			input[current] =temp
            current = j
		} else {
			break
		}
		
	}
	
  }
	 
	return input

}

func display(input []int) {
   
	for _,val := range input {
		fmt.Print(val," ")
	}

} 

func main () {


	var input []int
    var limit int
    fmt.Print("Enter the limit : ")
	fmt.Scan(&limit)
	fmt.Println("Enter the inputs : ")

    for i := 0; i < limit; i++ {
        var element int
		fmt.Scan(&element) 
		input = append(input, element)
		
	}
	    
	display(Sorting(input))
}