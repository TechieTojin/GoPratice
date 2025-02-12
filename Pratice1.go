package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int

	// Taking input from the user
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	fmt.Print("Enter third number: ")
	fmt.Scan(&num3)

	// Finding the largest number using conditional statements
	var largest int
	if num1 >= num2 && num1 >= num3 {
		largest = num1
	} else if num2 >= num1 && num2 >= num3 {
		largest = num2
	} else {
		largest = num3
	}

	// Printing the largest number
	fmt.Println("The largest number is:", largest)
}
