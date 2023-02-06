package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Welcome to Slices Program\n")

	var avengersList = []string{"Iron Main", "Hulk", "Winter Soldier"}
	fmt.Printf("\nType of avengersList is %T\n", avengersList)

	avengersList = append(avengersList, "Black Widow", "Captain America")

	fmt.Println("\nThe avengers list now looks like: ", avengersList, " and the length of the list is: ", len(avengersList))

	// Array Slicing
	// avengersList2 := append(avengersList[1:3])
	// fmt.Println(avengersList2)

	// Use about make keyword
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 145
	highScores[2] = 945
	highScores[3] = 7444

	fmt.Println("\nHigh scores list is: ", highScores, " and the length is: ", len(highScores))

	highScores = append(highScores, 123, 125, 124, 128, 256)
	fmt.Println("\nNow the Highscore slice looks like: ", highScores, " and the lenght is ", len(highScores))

	// Using the Sort package
	sort.Ints((highScores))
	fmt.Println("\nSorted array looks like: ", highScores)

	// Indexing
	fmt.Println("\nValue at index 2 in the slice is: ", highScores[2])

	highScores[2] = 12
	fmt.Println("\nNow the value at index 2 in the slice is: ", highScores[2])

	var courses = []string{"ReactJs", "NodeJs", "Swift", "Python", "Ruby"}
	fmt.Println("\nThe course slice looks like: ", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("\nSlice after removing second index looks like: ", courses)

}
