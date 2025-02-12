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

git init

git remote add origin https://github.com/TechieTojin/GoPratice.git

git remote -v

git add Pratice3.go

git commit -m "Added Pratice3.go"
git push -u origin main

git branch -M main
git push -u origin main
