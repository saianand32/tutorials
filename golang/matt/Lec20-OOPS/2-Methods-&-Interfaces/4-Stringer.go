package main

import "fmt"

// Student2 struct holds the name and marks of a student.
type Student2 struct {
	Name   string // Name of the student
	marks1 int    // Marks for subject 1
	marks2 int    // Marks for subject 2
}

// String method is part of the fmt.Stringer interface.
// It returns the difference between marks1 and marks2 as a string.
func (s Student2) String() string {
	return fmt.Sprint(s.marks1 - s.marks2)
}

// Number is a custom type alias for int.
type Number int

// String method for the Number type to implement the fmt.Stringer interface.
// It returns the string representation of the Number.
func (n Number) String() string {
	return fmt.Sprint(n)
}

// NotStringer is a custom type that does NOT implement fmt.Stringer.
type NotStringer int // Note: No String method here, so it doesn't implement fmt.Stringer.

func main() {
	// Example of types that implement the Stringer interface.

	var s fmt.Stringer

	// Student2 implements fmt.Stringer because it has a String() method.
	student := Student2{Name: "John", marks1: 90, marks2: 85}
	s = student    // No error, as Student2 implements Stringer
	fmt.Println(s) // Prints: 5 (difference between marks1 and marks2)

	// Number also implements fmt.Stringer because it has a String() method.
	var n Number = 90
	s = n          // No error, as Number implements Stringer
	fmt.Println(s) // Prints: 90

	// NotStringer does NOT implement fmt.Stringer because it lacks a String() method.
	var x NotStringer = 100
	// s = x // Error: cannot assign x (of type NotStringer) to s (of type fmt.Stringer)

	// We can use type assertion to check if x satisfies Stringer or not:
	// If x implements Stringer, s will be assigned the String method, otherwise it'll panic.
	if str, ok := any(x).(fmt.Stringer); ok {
		s = str
		fmt.Println(s)
	} else {
		fmt.Println("x does not implement fmt.Stringer")
	}
}
