package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{}

	slice = slices.Insert(slice, len(slice), 2)
	fmt.Println(slice)

}
