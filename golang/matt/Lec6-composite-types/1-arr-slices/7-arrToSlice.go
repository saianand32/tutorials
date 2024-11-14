package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}

	// convert array a to slice using [:]
	fmt.Printf("%T\n", a)    //[4]int array
	fmt.Printf("%T\n", b)    //[]int slice
	fmt.Printf("%T\n", a[:]) //[]int slice
}
