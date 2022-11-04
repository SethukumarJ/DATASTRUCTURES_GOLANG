package main


import (
	 "fmt"
	)

func Sorting (input []int) []int{
    
	for i := 0; i < len(input); i++ {
		
		for j := i; j < len(input); j++ {

			if input[i] > input[j] {

				temp := input[i]
				input[i] = input[j]
				input[j] = temp
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