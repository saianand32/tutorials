package main

import (
	"fmt"
)

func main() {
	a := 2

	fmt.Printf("%8T, %v \n", a, a)

	fmt.Printf("%8T, %[1]v \n", a)

	// ---there is no preincrement/decrement op
	// a++ --- correct
	// ++a --- sybtax err

	// both are same
	// 8 or any num can be used to give allign like a margin left while printing
	//[1] is the arg 1 in Printf function

	b := "sai"

	fmt.Scan(&b)
	fmt.Println(b)

}
