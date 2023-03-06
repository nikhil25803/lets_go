package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var signals = []string{
	"test",
}

// Usually pointers
var wg sync.WaitGroup

// Define Mutex, again a pointer in general
var mut sync.Mutex

func main() {
	// go greeter("Hello")
	// greeter("World")
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal("Error")
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Printf("%d Status Code for %s\n", res.StatusCode, endpoint)
}
