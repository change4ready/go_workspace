// Christian Tisby
// 1/4/2023
// Declaring and initializing multiple variables on one line

package main

import "fmt"

func main() {
	// initialize two string variables in the same statement
	var message, email string = "Hello, World!", "john@john.com"

	fmt.Println(message)
	fmt.Println(email)
}