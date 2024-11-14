package main

import "fmt"

// Declaring custom types (aliases) for various data types:

// Numbers is an alias for a slice of integers ([]int).
type Numbers []int

// Maps is an alias for a map with int keys and int values (map[int]int).
type Maps map[int]int

// Number is an alias for the int type.
type Number int

// Bool is an alias for the bool type.
type Bool bool

// Students is an empty struct type (no fields).
// Structs can be used to represent complex data types with multiple fields.
// An empty struct is often used as a placeholder or for signaling purposes.
type Students struct {
	// In this case, the Students struct is empty.
	// But in a real-world application, it might hold fields like Name, Age, etc.
}

func main() {
	// Declaring and initializing a variable of type Numbers (a slice of integers).
	// The slice is initialized with the elements 1, 2, 3, and 4.
	a := Numbers{1, 2, 3, 4} // This works, but it's a bit "awkward" because it's not as intuitive as declaring a slice of int.
	fmt.Println(a)           // Output: [1 2 3 4]

	// Declaring and initializing a variable of type Numbers using a slice literal.
	// This method is more direct and cleaner than the previous one.
	var b Numbers
	b = []int{1, 2, 3, 4} // Initialize b as a slice of integers with values 1, 2, 3, and 4.
	fmt.Println(b)        // Output: [1 2 3 4]

	// Declaring and initializing a variable of type Maps (a map from int to int).
	n := Maps{3: 30, 4: 40} // A map with keys 3 and 4, and values 30 and 40, respectively.
	fmt.Println(n)          // Output: map[3:30 4:40]

	// Declaring and initializing a variable of type Maps using a map literal.
	var m Maps
	m = map[int]int{1: 1} // A map with a single key-value pair: key 1 and value 1.
	fmt.Println(m)        // Output: map[1:1]

	// Declaring and initializing simple variables of custom types:
	var d Number = 89  // Declaring a variable of type Number (alias for int) and initializing it with the value 89.
	var e Bool = false // Declaring a variable of type Bool (alias for bool) and initializing it with the value false.
	fmt.Println(d, e)  // Output: 89 false

	// Declaring and initializing a variable of type Students (empty struct).
	// Since Students is an empty struct, we can't store any data in it,
	// but it can still be used to represent a type in your program.
	var student Students
	fmt.Println(student) // Output: {} (empty struct with no fields)

	// We can initialize an empty struct directly with shorthand notation:
	student2 := Students{}
	fmt.Println(student2) // Output: {} (another instance of an empty struct)
}
