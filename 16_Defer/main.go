package main

import "fmt"

func main() {
	fmt.Printf("Welcome to Defers in \nCalling the Greet function")
	greet()
	fmt.Printf("\nCalling the multiple functions with multiple defer:\n")
	multipleDefer()
}

func greet() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func multipleDefer() {
	defer fmt.Println("Line One")
	defer fmt.Println("Line Two")
	defer fmt.Println("Line Three")
	defer fmt.Println("Line Four")
	fmt.Println("Done")

}
