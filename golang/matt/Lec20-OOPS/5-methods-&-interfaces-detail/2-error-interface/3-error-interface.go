package main

import "fmt"

// errFoo is a custom error type that implements the error interface
type errFoo struct {
	err  error  // the actual error
	path string // some path info
}

// Implement the Error() method for errFoo to satisfy the error interface
func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}

// XYZ returns a pointer to errFoo (a concrete type), which might be nil.
func XYZ(a int) *errFoo {
	// Return a nil pointer (concrete type is *errFoo, value is nil)
	return nil
}

// XYZ2 returns an error interface, which can also be nil.
func XYZ2(a int) error {
	// Return a nil error (interface is nil here)
	return nil
}

func main() {
	// Example 1: Using XYZ function
	var err error = XYZ(1) // BAD: Interface gets a nil concrete pointer (type *errFoo, value nil)

	// The interface holds a non-nil type (*errFoo) but a nil value
	// So the comparison with nil is true in Go, and the condition will evaluate to true
	if err != nil {
		// Even though XYZ returned a nil pointer, the interface is not nil
		// The interface holds a *errFoo type but the value is nil
		fmt.Println("oops") // Output: "oops" (because the interface is not nil)
	}

	// Example 2: Using XYZ2 function
	var err2 error = XYZ2(1) // Interface gets a nil concrete pointer, the error interface is truly nil
	// In this case, XYZ2 returns a truly nil error interface

	if err2 != nil {
		// Here the error interface itself is nil, so this condition is false
		fmt.Println("oops") // This does not get printed
	}

	// NOTE: Always declare error like this (var err error) and dont do (err := nil) as u need
	// to specify that this err is of type error interface

	// FUN_TRY: Try to find out how to find if interface is nil or interface has a nil value
}
