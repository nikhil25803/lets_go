package main

import "fmt"

// Declare some constants value
const LoginToken string = "This is a token"

func main() {
	// Strings type variables
	var username string = "Nikhil"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// Boolean type values
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// Uint 8 type integrer
	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type: %T \n", smallInt)

	// Float datatypes
	var smallfFloat float64 = 255.0131564651315
	fmt.Println(smallfFloat)
	fmt.Printf("Variable is of type: %T \n", smallfFloat)

	// Default values and Alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// Another way of declaring a varibales
	var website = "github.com/nikhil25803"
	fmt.Println(website)

	// No var keyword
	numberOfUser := 300000
	fmt.Println(numberOfUser)
	fmt.Printf("This var is of type: %T \n", numberOfUser)

	// Accessing the Global variable
	fmt.Println(LoginToken)
}
