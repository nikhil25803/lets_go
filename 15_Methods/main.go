package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in Golang")
	// No inheritance in Go lang, also there is no super and parent
	nikhil := User{"Nikhil", "nikhil25803@gmail.com", true, 18}
	fmt.Println("The struct looks like: ", nikhil)
	fmt.Printf("Nikihl details are: %+v\n", nikhil)
	fmt.Printf("Name is: %v and Email is: %v", nikhil.Name, nikhil.Email)
	nikhil.GetStatus()
	nikhil.NewMail()
	fmt.Printf("\nName is: %v and Email is: %v", nikhil.Name, nikhil.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Functions for Structs are called Methods
func (u User) GetStatus() {
	fmt.Println("\nIs user active? ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@mail.dev"
	fmt.Printf("Email of %v user is %v", u.Name, u.Email)
}
