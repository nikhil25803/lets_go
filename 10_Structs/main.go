package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in Golang")
	// No inheritance in Go lang, also there is no super and parent
	nikhil := User{"Nikhil", "nikhil25803@gmail.com", true, 18}
	fmt.Println("The struct looks like: ", nikhil)
	fmt.Printf("Nikihl details are: %+v\n", nikhil)
	fmt.Printf("Name is: %v and Email is: %v", nikhil.Name, nikhil.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
