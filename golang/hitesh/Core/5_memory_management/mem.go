package main

import "fmt"

func main() {
	// 2 ways

	// 1. new()
	// -- Allocate memory but no initialization
	// -- u will get a memory address gets back a pointer
	// -- zeroed storage(no data initially)

	// 2. make()
	// -- majority times we use this
	// -- Allocate memory but and initialization
	// -- u will get a memory address
	// -- non zeroed storage

	// make is applicable only for slices, maps and channels
	//new is for all types

	// 3. usage of make
	map1 := make(map[int]string)
	map1[1] = "sai"
	fmt.Println(map1) //map[1:sai]

	// 4. usage of new
	age := new(int)
	*age = 22
	fmt.Println(age)  //0xc000096088
	fmt.Println(*age) //22
}
