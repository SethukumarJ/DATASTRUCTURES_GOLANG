package main

import (
	"fmt"
	"strings"
)

var frequency = make(map[string]int, 3)

func Frequency(date string, value string){

	var date1,f string
	fmt.Println("enter the date : ")
	fmt.Scan(&date1)
	fmt.Println("enter the frequency : ")
	fmt.Scan(&f)


	s := strings.Split(date1,"")


	day := []int{}
	mnth := []int{}

	for i, _ := range s[0]{

		day = append(day,int(byte(s[0][i])))
	}

	for i, _ := range s[1]{

		mnth = append(mnth,int(byte(s[1][i])))
	}
	
if value == "month" {

		if len(mnth) ==2 {

			mnth[1]++

		} else {
			mnth[0]++
		}
	}
	if value == "day" {

		if len(mnth) ==2 {

			day[1]++

		} else {
			day[0]++
		}
	}
	if value == "week" {

		if len(day) ==2 {

			day[1]  = day[1] + 7

		} else {
			day[0]++
		}
	}


	
}

func main() {

	frequency["month"] = 1
	frequency["week"] = 7
	frequency["day"] = 30
	Frequency("10/11/2002", "month") 
}
