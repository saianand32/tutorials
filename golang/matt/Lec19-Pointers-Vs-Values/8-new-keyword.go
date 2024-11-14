package main

import "fmt"

// Note prefer not to use new as its nothing special
// using new(type) is same as &Type{}

//say new(Person) and &Person{} do similar things infact & one is better
// So always prefer & method

//Also new does not mean allocation on heap in go... in java new always allocates on heap but
// in go escape analysis is done and that determines stack or heap allocation

// Define a struct type
type Person struct {
	Name string
	Age  int
}

func main() {
	// Example 1: Using `new` to allocate memory for a simple type (int)
	fmt.Println("Example 1: Using new with int")
	ptrInt := new(int)                       // Creates a pointer to an int, initialized to 0
	fmt.Println("Value of ptrInt:", *ptrInt) // Dereferencing the pointer
	*ptrInt = 42
	fmt.Println("Updated value of ptrInt:", *ptrInt)

	// Example 2: Using `new` to allocate memory for a struct
	fmt.Println("\nExample 2: Using new with a struct")
	ptrPerson := new(Person)                         // Creates a pointer to a Person struct, with zero values for fields
	fmt.Println("Before initialization:", ptrPerson) // Prints zero value fields: Name="", Age=0
	ptrPerson.Name = "John"
	ptrPerson.Age = 30
	fmt.Println("After initialization:", ptrPerson)

	// Example 3: Using `new` to allocate memory for an array
	fmt.Println("\nExample 3: Using new with an array")
	ptrArray := new([3]int) // Creates a pointer to an array of 3 integers, initialized to [0, 0, 0]
	fmt.Println("Before modification:", *ptrArray)
	ptrArray[0] = 10
	ptrArray[1] = 20
	ptrArray[2] = 30
	fmt.Println("After modification:", *ptrArray)

	// Example 4: Using `new` with a slice
	fmt.Println("\nExample 4: Using new with a slice")
	ptrSlice := new([]int)                 // Creates a pointer to a slice of ints, initialized to nil
	*ptrSlice = append(*ptrSlice, 1, 2, 3) // Dereferencing the pointer and appending values
	fmt.Println("Slice values:", *ptrSlice)
}
