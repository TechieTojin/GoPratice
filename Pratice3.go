package main

import "fmt"

func main() {
	var year int

	// Take user input
	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	// Check leap year condition
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println(year, "is a leap year.")
	} else {
		fmt.Println(year, "is not a leap year.")
	}
}

