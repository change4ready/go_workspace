// Christian Tisby
// 1/4/2023
// Variable Scope - Global Variable Scope
package main

import "fmt"

// global variable
var message string = "Hello, World!"

func main() {
	// local variable
	var email string = "john@john.com"

	fmt.Println(message)
	fmt.Println(email)
}