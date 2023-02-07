package main

import "fmt"

func main() {
	fmt.Println("Welcome to Conditional Statements Program")

	loginCount := 100

	var result string

	if loginCount < 10 {
		result = "If Statement"
	} else if loginCount >= 10 && loginCount < 20 {
		result = "Else if Statement"
	} else {
		fmt.Println("Else statemet")
	}

	fmt.Println(result)

}
