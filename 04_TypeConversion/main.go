package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the type conversion program")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	input_name, _ := reader.ReadString('\n')
	fmt.Print("Enter your age: ")
	input_age, _ := reader.ReadString('\n')
	fmt.Print("Enter your salary: ")
	input_salary, _ := reader.ReadString('\n')

	fmt.Print("All  details are given in Strings.\nAfter typeconversion, the datatypes f each details are: \n")
	age, _ := strconv.ParseInt(strings.TrimSpace(input_age), 10, 64)
	salary, err := strconv.ParseFloat(strings.TrimSpace(input_salary), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type of name is: %T\n", input_name)
		fmt.Printf("Data type of age is: %T\n", age)
		fmt.Printf("Data type of salary is: %T\n", salary)
	}

}
