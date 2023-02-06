package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study program")

	presentTime := time.Now()
	fmt.Print(presentTime)

	// This outputs - 2023-02-06 22:05:12.8190491 +0530 IST m=+0.001639001

	// We always need to give the date and time in this format only i.e Month-Date-Year Hour:Minute:Second Day
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Print('\n')

	createDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
