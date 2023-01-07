// Christian Tisby
// 1/5/2023
// Attempting to add two values

package main

import "fmt"

func main() {
	var firstNumber string
	var secondNumber string

	fmt.Print("Enter the first integer: ") // user prompt
	fmt.Scan(&firstNumber) // store input

	fmt.Print("Enter the second integer: ") // user prompt
	fmt.Scan(&secondNumber)

	fmt.Println(firstNumber + secondNumber) // addition of two strings
}