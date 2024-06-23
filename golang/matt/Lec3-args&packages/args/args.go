package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide args")
	}

	strNum := os.Args[1]
	// these args are strings

	num, err := strconv.Atoi(strNum)

	if err == nil {
		fmt.Println("Int arg =", num)
	}
}
