package main

func main() {
	// A struct with no fields is actually usefull

	// 1. making a set with using map with empty struct instead of bool to take less space

	// a- using make
	set := make(map[int]struct{})
	set[1] = struct{}{}
	//b- using map literal
	set2 := map[int]struct{}{}
	set2[1] = struct{}{}

	// IMPORTANT: empty struct{} is a SINGLETON -- any empty struct points to this in memory.const
	// cannot have multipleInstances of struct{} in diff locations.
	// Infact go stores this in a particular mem location.
	// and even an empty slice points to this struct{} thats how we differentiate an empty slice and nil slice.
}
