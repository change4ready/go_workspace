// Christian Tisby
// 1/6/2023
// Exercise 3.2 Creating variables
// Write a program that includes my first nam, street address, and the year i was born

package main

import "fmt"

func main() {
	var firstName string = "Chris"
	var streetAddress string = "12589 Lake Madison Ln"
	var yearBorn int = 1982

	fmt.Println(firstName)
	fmt.Println(streetAddress)
	fmt.Println(yearBorn)
}