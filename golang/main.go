package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument")
		return
	}

	strNum := os.Args[1] // The first argument

	// Command-line arguments are always gotten as strings
	num, err := strconv.Atoi(strNum) // Convert the string argument to an integer

	if err != nil {
		fmt.Println("couldnt convert to int")
	}

	fmt.Printf("Number argument = %d", num)
}

// go run main.go 12
// output --> Number argument = 12
