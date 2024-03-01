package main

import (
	"errors"
	"fmt"
)

// 1. Named returns aare always better  to use and below shows example of errors throwing
// Guard Clause - is an early return from a function as shown below

func divide(a float64, b float64) (ans float64, err error) {

	if b == 0.0 {
		return -1.0, errors.New("Cant divide by 0")
	}
	return a / b, nil
}

func main() {
	ans, err := divide(2, 0)

	fmt.Println(ans)
	fmt.Println(err)

	// 2. Anonymous functions and IIFE

	add := func(a int, b int) (sum int) { //anonymous function
		return a + b
	}

	fmt.Println(add(2, 4))

	func() {
		fmt.Println("Iam IIFE")
	}()

}
