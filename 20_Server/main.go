package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web requests in Golang")
	PerformGetRequest()
	print("\n")
	PerformPostRequest()
	print("\n")
	PerformFormRequest()

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

func PerformPostRequest() {
	myurl := "http://localhost:3000/post"

	requestBody := strings.NewReader(`
		{
			"Course":"Lets Go with Golang",
			"Price":0,
			"Platform":"YouTube"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))
}

func PerformFormRequest() {
	myurl := "http://localhost:3000/postForm"

	data := url.Values{}
	data.Add("firstName", "Nikhil")
	data.Add("lastName", "Raj")
	data.Add("email", "nikhil25803@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
