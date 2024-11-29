package main

import "fmt"

// PairFiz represents a rectangle with length and breadth
type PairFiz struct {
	len     int
	breadth int
}

// String method for PairFiz to control how it is printed
func (p PairFiz) String() string {
	return fmt.Sprintf("length = %d, breadth = %d", p.len, p.breadth)
}

// FizGiz is a composite struct that includes PairFiz and an additional boolean field Fiz
type FizGiz struct {
	*PairFiz // Embedding PairFiz as a pointer
	Fiz      bool
}

func main() {
	// Create an instance of FizGiz with embedded PairFiz
	a := FizGiz{
		PairFiz: &PairFiz{len: 2, breadth: 4}, // Initialize PairFiz inside FizGiz
		Fiz:     true,
	}

	// Print the FizGiz instance
	// Note: Since PairFiz has a String method, it gets automatically called
	fmt.Println(a) // Output: length = 2, breadth = 4
}

// here u see that PairFiz String method got called this shows that
// even the methods of the embedded types get promoted here not only variables of embedded types
