package main

import "fmt"

// Student struct holds the name and marks of a student.
type Student struct {
	Name   string // Name of the student
	marks1 int    // Marks for subject 1
	marks2 int    // Marks for subject 2
}

// String method is part of the fmt.Stringer interface.
// It allows us to define a custom string representation for the Student struct.
// In this case, it returns the difference between marks1 and marks2 as a string.
func (s Student) String() string {
	return fmt.Sprint(s.marks1 - s.marks2) // Returns the difference of marks1 and marks2
}

func main() {
	// In Go, the String() method allows us to control how an instance of a struct is printed when used with functions like fmt.Println.
	// It behaves somewhat like the toString method in Java or the __str__ dunder method in Python.

	// Example usage:
	student := Student{
		Name:   "Sai", // Name of the student
		marks1: 95,    // Marks in subject 1
		marks2: 88,    // Marks in subject 2
	}

	// When we print the student, the custom String() method is called, and it prints the difference between marks1 and marks2.
	// Output: 7
	fmt.Println(student) // This will print: 7
}
