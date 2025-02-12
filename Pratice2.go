package main

import "fmt"

func main() {
	// Predefined target number
	target := 42
	var guess int

	fmt.Println("Guess the number (between 1 and 100):")

	for {
		// Take user input
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)

		// Check the guess
		if guess < target {
			fmt.Println("Too low. Try again.")
		} else if guess > target {
			fmt.Println("Too high. Try again.")
		} else {
			fmt.Println("Correct! You guessed the number:", target)
			break
		}
	}
}
