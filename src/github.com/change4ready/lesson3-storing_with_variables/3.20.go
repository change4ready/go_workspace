// Christian Tisby
// 1/6/2023
// Arcitecture-independent integers
package main

import "fmt"

func main() {
	var age int8 = 20 // signed 8-bit integer
	var port int16 = 80 // signed 16-bit integer
	var zipcode int32 = 90000 // signed 32-bit integer
	var phone int64 = 7322335624 // signed 64-bit integer
	var phone2 uint64 = 7322335624 // unsigned 64-bit integer
	var score int64 = -1 // signed 64-bit integer w/ negative value
	// var score uint64 = -1 <- this is an illegal variable

	fmt.Println("age int8", age)
	fmt.Println("port int16", port)
	fmt.Println("zipcode int32", zipcode)
	fmt.Println("phone int64", phone)
	fmt.Println("phone2 unit64", phone2)
	fmt.Println("score int64", score)

}