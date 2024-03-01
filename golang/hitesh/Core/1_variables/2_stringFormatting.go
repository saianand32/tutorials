package main

import "fmt"

func main() {

	// mainly 2 ways Printf and Sprintf

	// 1. Default %v for an kind of variables

	fmt.Printf("I am %v years old", 10)

	str := fmt.Sprintf("I am %v years old", "way too many")
	fmt.Println(str)

	i := 90
	b := true
	c := 8 + 5i

	fmt.Printf("int i is : %v \n", i)
	fmt.Printf("bool b is : %v \n", b)
	fmt.Printf("complex c is : %v \n", c)

	//2. %d for int
	//3. %f for float and %.2 f for decimals upto 2
	fmt.Printf("float upto 3 decimals : %.3f \n", 8.9978)

	// 4. %c for char(there is no char type use rune)
	// 5. %t for bool
	// 6. %s for string
	// 7. for complex just use %v

	// print Type of a variable (use %T to print type of variable)
	fmt.Printf("%T", c) //complex128
	fmt.Printf("%T", b) //bool
	fmt.Printf("%T", i) //int

}
