package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files in Golang.")
	content := "This is a test content to be written in a txt file"

	file, err := os.Create("./myGofile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./myGofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside file is \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
