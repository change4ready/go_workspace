// Christian Tisby
// 1/4/2023
// Example of using an invalid variable name

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