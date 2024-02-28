package main

import "fmt"

func main() {

	// Using the fmt package itself

	var num int
	var name string
	var flag bool

	fmt.Scanln(&num)
	fmt.Scanln(&name)
	fmt.Scanln(&flag)

	fmt.Printf("%T, %T, %T", num, name, flag)

}
