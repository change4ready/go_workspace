// Christian Tisby
// 1/5/2023
// Trying to store a float in an integer
package main

import "fmt"

func main() {
	quantity := 100 // assigns int type to variable
	fmt.Printf("quantity type: %T\n", quantity)

	// quantity = 99.5 // store float in same variable: error
	//fmt.Println(quantity)
}