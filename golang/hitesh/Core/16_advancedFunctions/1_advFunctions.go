package main

import "fmt"

//functions can be passed like variables

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func aggregate(a, b int, arithmetic func(int, int) int) int {
	return arithmetic(a, b)
}

func main() {
	a, b := 10, 20

	fmt.Println(aggregate(a, b, add)) //30
	fmt.Println(aggregate(a, b, mul)) //200
}

//first class function = simple function
//Higher order functions = A function that takes another first class func as arg
