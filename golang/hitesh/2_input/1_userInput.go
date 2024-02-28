package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// ******* 1. Using bufio *********
	reader := bufio.NewReader(os.Stdin)

	//this is a comma ok syntax
	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(str)
	}

	fmt.Println()

	//taking number input

	age, err := reader.ReadString('\n')

	fmt.Println(age)

	//need to trim the spaces and the extra \n we took input
	numAge, err := strconv.ParseFloat(strings.TrimSpace(age), 64)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("%T %T \n", numAge, age) // float64 string
	fmt.Println(numAge + 1)

	// similarly we have
	// strconv.ParseFloat, ParseInt, ParseUint, ParseBool, ParseComplex
}
