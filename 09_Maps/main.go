package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps Program in Go lang")

	language := make(map[string]string)
	language["JS"] = "Javascript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"
	language["Go"] = "Golang"

	fmt.Println("The language map looks like: \n", language)
	fmt.Println("JS shorts for: ", language["JS"])

	delete(language, "RB")
	fmt.Println("Now the language map looks like: \n", language)

	// Loop on a Map
	for key, value := range language {
		fmt.Printf("For key %v, value is: %v\n", key, value)
	}

}
