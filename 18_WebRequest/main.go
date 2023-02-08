package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/"

func main() {
	fmt.Println("Welcomme to Web Request program")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T", response)

	defer response.Body.Close()

	// To read the repsonse

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(body)
	fmt.Println(content)
}
