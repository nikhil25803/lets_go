package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to Functions Program")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	// slice := []int{1, 2, 3, 4, 5, 6}
	proRes := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Sum of the slice is: ", proRes)

}

func adder(valOne int, val2 int) int {
	return valOne + val2
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Hello from Golang")
}
