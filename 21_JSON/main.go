package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Working with JSON in Golang")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCOurses := []courses{
		{"ReactJS", 299, "learncodeonline.com", "abc123", []string{"web-dev", "js"}},
		{"MERN", 199, "learncodeonline.com", "abc123", []string{"web-dev", "js"}},
		{"FARM", 499, "learncodeonline.com", "abc123", nil},
	}

	// Package this data as JSON
	finalJson, _ := json.MarshalIndent(lcoCOurses, "", "\t")
	fmt.Println(string(finalJson))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN",
		"Price": 199,
		"Website": "learncodeonline.com",
		"tags": [
				"web-dev",
				"js"
			]
	}
	`)
	var lcoCOurses courses
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCOurses)
		fmt.Printf("%#v\n", lcoCOurses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// There are some cases where we want to add only data to key and value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}

}
