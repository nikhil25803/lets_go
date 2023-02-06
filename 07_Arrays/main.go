package main

import "fmt"

func main() {
	fmt.Printf("Welcome to Array program\n")

	var avengersList [4]string

	avengersList[0] = "Iron Man"
	avengersList[1] = "Hulk"
	avengersList[3] = "Captain America"

	fmt.Println("Avengers List list is: ", avengersList)
	fmt.Println("Length of the list is: ", len(avengersList))

	var dcList = [5]string{"Batman", "Superman", "Aquaman"}
	fmt.Println("The DC list is: ", dcList)
	fmt.Println("Length of the list is: ", len(dcList))
}
