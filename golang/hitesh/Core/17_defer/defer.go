package main

import "fmt"

func main() {
	// the comments are the order of printing
	fmt.Println("hello") //1
	fmt.Println("world") //2

	defer fmt.Println("hello") //6
	defer fmt.Println("world") //5
	defer fmt.Println("anand") //4
	fmt.Println("sai")         //3

	//defer makes those lines execute in LIFO order
}
