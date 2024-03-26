package main

import "fmt"

func main() {

	// Note : the below methods can be used to scan almost all but cant scan string with spaces u need to use bufio

	// Using the fmt package itself

	// 1. The Scan() function receives the user input in raw format as space-separated values and stores them in the variables. Newlines count as spaces.

	var i, j int

	fmt.Print("Type two numbers: ")
	fmt.Scan(&i, &j)
	fmt.Println("Your numbers are:", i, "and", j)

	// 2. The Scanln() function is similar to Scan(), but it stops scanning for inputs at a newline (at the press of the Enter key.).

	var num int
	var name string
	var flag bool

	fmt.Scan(&num)
	fmt.Scanln(&name)
	fmt.Scanln(&flag)

	fmt.Printf("%T, %T, %T", num, name, flag)

	// 3. The Scanf() function receives the inputs and stores them based on the determined formats for its arguments.

	var a, b int
	fmt.Print("Type two numbers: ")
	fmt.Scanf("%v %v", &a, &b)
	fmt.Println("Your numbers are:", i, "and", j)

}
