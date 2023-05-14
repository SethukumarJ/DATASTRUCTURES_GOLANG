package main

import "fmt"

func main() {

	// a := make(map[string][]int)

	// a["first"] = []int{1, 3, 4, 5}
	// a["second"] = []int{1, 9, 4, 5}
	// a["third"] = []int{1, 6, 4, 5}

	// for r, v := range a {

	// 	fmt.Println(r,v)
	// }

	a := input()
	
	fmt.Println(a)
}

func input() map[string][]int {

	a := make(map[string][]int)

	var length int
	var key string
	var value []int
	var V int
	var SizeOfArray int
	
	
	fmt.Println("enter the length :")
	fmt.Scan(&length)

	for i := 1; i <= length; i++ {

		fmt.Println("enter teh key: ")
		fmt.Scan(&key)
		fmt.Println("enter the size of the array")
		fmt.Scan(&SizeOfArray)

		for j:=1; j<=SizeOfArray; j++  {

			fmt.Println("Enter the values")

			fmt.Scan(&V)

			
			value = append(value, V)
	

			}
		
			
		a[key] = value
		value = []int{}
		
	}
	return a

}
