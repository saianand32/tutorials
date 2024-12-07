package main

import "fmt"

//These are interfaces in fmt
//type Formatter
// type GoStringer
// type ScanState
// type Scanner
// type State
// type Stringer

//but we will look into 2 of them
// 1. Stringer
// 2. GoStringer

type MyStruct struct {
	Name string
	Age  int
}

// Implement the Stringer interface for MyStruct
func (m MyStruct) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", m.Name, m.Age)
}

// Implement the GoStringer interface for MyStruct
func (m MyStruct) GoString() string {
	return fmt.Sprintf("MyStruct{Name: %q, Age: %d}", m.Name, m.Age)
}

func main() {

	m := MyStruct{"Alice", 30}

	// 1. ****** Stringer **********
	// type Stringer interface {
	// 	String() string
	// }
	// Stringer is an interface in the fmt package that defines a method to convert an object to a string.
	// The Stringer interface is primarily used to define how types should be represented as a string,
	// especially when you call fmt.Println, fmt.String or %s formatting verb.
	// If your type implements the Stringer interface, you can control how it is printed when used with formatting functions.

	fmt.Println(m) // Output: Name: Alice, Age: 30

	// 2. GoStringer
	// type GoStringer interface {
	// 	GoString() string
	// }
	// GoStringer is an interface that allows types to define how they should be represented as Go source code.
	// It is commonly used with the fmt package's %#v formatting verb, which attempts to provide a Go-syntax
	// representation of the object.

	fmt.Printf("%#v\n", m) // Output: MyStruct{Name: "Alice", Age: 30}
}
