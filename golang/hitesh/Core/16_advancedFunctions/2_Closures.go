package main

import "fmt"

//1. Closures def same as js

func fun() func(int) int {
	a := 10
	return func(b int) int {
		return a + b
	}
}

func main() {
	fun2 := fun()

	fmt.Println(fun2(9)) //19 even though fun2 doesnt have any knolege of a=10  it still maintaines access to lexical scope
}
