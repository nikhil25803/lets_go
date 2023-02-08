package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web requests in Golang")
	PerformGetRequest()

}

func PerformGetRequest() {
	myurl := "http://localhost:3000/"
	res, err := http.Get(myurl)
	if err != nil {
		panic((err))
	}

	defer res.Body.Close()
	fmt.Println("Status code is: ", res.Status)

	var responsestring strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	bytecount, _ := responsestring.Write(content)

	fmt.Println("Bytecount is: ", bytecount)
	fmt.Println("Raw data is is: ", responsestring.String())

}
