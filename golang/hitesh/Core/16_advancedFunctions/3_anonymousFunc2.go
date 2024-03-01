package main

import "fmt"

func fun(msg string, fun2 func(int, int) int) {
	fmt.Println(fun2(2, 3))
	fmt.Println(msg)
}

func main() {
	//calling fun function anonymous like a callback function

	fun("add", func(a int, b int) int {
		return a + b
	}) //5 add
}
