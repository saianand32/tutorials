package main

import "fmt"

// Define a type IntSet
type IntSet struct{}

// Define a method String() with a pointer receiver
func (*IntSet) String() string {
	return ""
}

func main() {
	// Error: `String()` method needs a pointer receiver
	// here IntSet{} is a literal it doesnt have an address so we cant have a pointer to a literal ... its not tied any where
	var a = IntSet{}.String() // error String needs *IntSet

	// OK: `String()` method can be called on a variable of type IntSet.. automatically uses &s
	var s IntSet
	var z = s.String() // ok

	// OK: `*IntSet` (pointer) implements fmt.Stringer, so we can assign a pointer to `fmt.Stringer`
	var _ fmt.Stringer = &s // ok

	// Error: `IntSet` (value) does not implement fmt.Stringer because String() requires a pointer receiver
	var _ fmt.Stringer = s // error no String method
}
