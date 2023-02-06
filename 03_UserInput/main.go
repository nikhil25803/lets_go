package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your name: ")

	// Comma Ok or Error Ok Synatx
	input, _ := reader.ReadString('\n')
	fmt.Print("Hey, ", input, "welcome to the program")
}
