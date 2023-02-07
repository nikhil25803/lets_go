package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Golang")

	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	// For Loops
	fmt.Printf("\nFor Loop 1\n")
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Printf("\nFor Loop 2\n")
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Printf("\nFor Loop 3\n")
	for index, day := range days {
		fmt.Printf("Index is %v and value is: %v\n", index, day)
	}

	fmt.Printf("\nFor Loop 4\n")
	rogueValue := 1
	for rogueValue < 10 {
		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

	// Loope With break and Continue
	fmt.Printf("\nFor Loop 5\n")
	value := 1
	for value < 10 {
		if value < 5 {
			fmt.Println("Value is: ", value)
		} else if value == 5 {
			value++
			continue
		} else if value == 6 {
			// Go to also works as a break statement
			goto line
		} else if value == 8 {
			break
		} else {
			fmt.Println("Value is: ", value)
		}
		value++
	}

line:
	fmt.Println("Goto value is: 6")
}
