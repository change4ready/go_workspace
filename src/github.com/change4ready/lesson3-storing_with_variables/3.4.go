// Christian Tisby
// 1/4/2023
// Declaring multiple variables in one statement

package main

import "fmt"

func main() {
	var message, email string // declare two string variables
	message = "Hello, World!" // assign a variable to one variable
	email = "john@john.com" // assign a value to the second variable

	fmt.Println(message)
	fmt.Println(email)
}