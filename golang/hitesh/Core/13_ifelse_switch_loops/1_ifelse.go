package main

import "fmt"

func main() {

	// 1. if else

	num := 20

	if num > 50 {
		fmt.Println("greater than 50")
	} else if num < 50 && num > 20 {
		fmt.Println("lesser than 50 greater than 20")
	} else {
		fmt.Println("lesser than 20")
	}

	//2. if else with direct values

	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// 3. if else with short statement

	if age := 21; age > 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("not adult")
	}

	if age, name := 21, "John"; age > 18 {
		fmt.Println(name, "is an adult")
	} else {
		fmt.Println(name, "is not an adult")
	}

}
