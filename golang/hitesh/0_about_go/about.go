package main

import "fmt"

func main() {

	// Gophers (ppl building with golang)

	// 1. Go can run files directly without VM
	// 2. But executables are os specific

	// Golang doesnt have trycatch

	// is it oop language- Not exactly no classes, method overriding etc nothing
	// but has structs to achieve similar things

	// lexer does a lot of work(it puts semicolons, imports modules automatically etc)

	// To create a go mod file
	// $ go mod init module_name

	// To get help
	// $ go help
	// $ go help command_name

	// to run a go file
	// $ go run prac.go

	//Printing in go
	fmt.Println("Hello world")
	str := fmt.Sprintln("assign to string")
	fmt.Println(str)
	// also can use Printf and Sprintf

}
