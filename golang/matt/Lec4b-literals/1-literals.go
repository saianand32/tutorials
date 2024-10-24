package main

import "fmt"

func main() {
	// In Go, literals are constant values that are directly assigned to variables. 
	// They represent fixed values of basic data types and can be written directly into your code. 
	// Go supports several types of literals, including integer, floating-point, boolean, string, rune, and complex literals.

	// int literals
	a := 125
	b := 1_32_000 //same as 1,32,000 made easy to understand but instead of comma we use _

	// float literals
	c := 123.99
	d := 1_23.67
	e := 1_23.67e3

	// bool literals
	f := true
	g := false

	// str literals
	h := "john"

	// complex literals
	i := 7 + 4i
	j := 70_000 + 8_000i

	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}
