package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in golang")

	// current time
	presentTime := time.Now()
	fmt.Println(presentTime)

	// always need to use Jan 2, 2006 as a date for formatting
	// time will always need to be 15:04:05
	fmt.Println(presentTime.Format("01/02/2006 Monday, 15:04:05"))

	// creating time from values
	createdDate := time.Date(2020, time.April, 20, 4, 20, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
