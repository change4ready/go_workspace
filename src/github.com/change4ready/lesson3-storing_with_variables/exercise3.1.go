// Christian Tisby
// 1/6/2023
// Exercise 1 Fixing Problems

package main

import "fmt"

func main() {
	// the variable declaration below is invalid 
	// due to the invalid naming convention
	//var 0_email string = "john@john.com"

	// I could fix it by simply removing the 0_ and just name the variable email
	var email string = "john@john.com"
	fmt.Println(email)
}