package main

import "fmt"

// 1. Variadic functions (or) Varargs
func fun(age int, names ...string) {
	fmt.Println(age)
	fmt.Println(names[2])
}

func main() {
	fun(2, "sai", "anand", "ankur", "saishwar") //anand

	//2. Spread operator (one use case is to spread the slice and pass into as args of a function)
	arr := []string{"sai", "ankur", "annand"}
	fun(22, arr...) //annand
}
