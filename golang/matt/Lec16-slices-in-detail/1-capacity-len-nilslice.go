package main

import "fmt"

func main() {

	// 1. nil slice
	var s []int

	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil) // 0, 0, []int,  true, []int(nil)

	// 2. empty slice
	t := []int{}
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil) // 0, 0, []int, false, []int{}

	// 3. slice with using make with one param
	u := make([]int, 5)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil) // 5, 5, []int, false, []int{0, 0, 0, 0, 0}

	// 4. slice with using make with two params
	v := make([]int, 0, 20)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil) // 5, 20, []int, false, []int{}

	// see the next image for image description
}
