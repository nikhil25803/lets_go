package main

import "fmt"

func main() {
	fmt.Printf("Welcome to Pointers program \n")

	// Declaring a pointer
	var ptr *int
	fmt.Println("Value of pointer is ", ptr)

	// Another symbol
	myNumber := 25
	var ptr2 = &myNumber
	fmt.Println("value of the pointer with value is: ", ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println("Value of the first pointer now is: ", myNumber)

}
