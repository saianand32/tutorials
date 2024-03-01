package main

import (
	"fmt"
	"time"
)

func main() {

	// 1. current time
	presentTime := time.Now()
	// fmt.Println(presentTime)

	// 2. using   Format
	fmt.Println(presentTime.Format("01-02-2006"))                 // remember this exact date u have to put no other (mm/dd/yyyy)
	fmt.Println(presentTime.Format("01/02/2006 Monday"))          //have to use Monday to print in this format M is capital
	fmt.Println(presentTime.Format("01/02/2006 15:04:05 Monday")) // have to use 15:04:05 to print the timings in (hh/mm/ss)

	// Creating a date

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01/02/2006 15:04:05 Monday"))

}
