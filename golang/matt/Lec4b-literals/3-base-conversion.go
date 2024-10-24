package main

import (
	"fmt"
)

func main() {
	a := 129

	// 1. Binary representation
	binary := fmt.Sprintf("%b", a)
	fmt.Println("Binary:", binary) //Binary: 10000001

	// 2. Octal representation
	octal := fmt.Sprintf("%o", a)
	fmt.Println("Octal:", octal) //Octal: 201

	// 3. Hexadecimal representation
	hexadecimal := fmt.Sprintf("%x", a)
	fmt.Println("Hexadecimal:", hexadecimal) //Hexadecimal: 81

	//cannot use boolean here unlike python.

}
