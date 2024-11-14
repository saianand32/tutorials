package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	// 1. append
	a = append(a, 6)
	fmt.Println(a) //[1 2 3 4 5 6]

	// 2. append 2 slices
	b := []int{7, 8, 9, 10}
	a = append(a, b...)
	fmt.Println(b) // [1 2 3 4 5 6 7 8 9 10]

	// 3. copy
	c := make([]int, len(a)) //need to fix lenght same as a else some will be missed out
	copy(c, a)
	fmt.Println(c) // [1 2 3 4 5 6 7 8 9 10]
}
