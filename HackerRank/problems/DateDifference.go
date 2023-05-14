package main

import (
	"fmt"
	"time"
)

func main() {

	DateCalculate()

}

func DateCalculate() {

	var fd string
	var sd string
	fmt.Println("Enter the first date : ")
	fmt.Scan(&fd)
	fmt.Println("Enter the last date : ")
	fmt.Scan(&sd)


	firstDate := time.Date(2022, 4, 13, 1, 0, 0, 0, time.UTC)
    secondDate := time.Date(2021, 2, 12, 5, 0, 0, 0, time.UTC)

	fmt.Println(firstDate)
    difference := firstDate.Sub(secondDate)
	
	fmt.Printf("Days: %d\n", int64(difference.Hours()/24))
}
