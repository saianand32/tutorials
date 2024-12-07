package main

import (
	"fmt"
)

// Add is a simple function that adds two integers.
func Add(n1, n2 int) int {
	return n1 + n2
}

// Point represents a 2D point in the plane.
type Point struct {
	X float64
	Y float64
}

// Method with a value receiver: Distance calculates a "distance" based on both points' coordinates.
// This method operates on a copy of the Point.
func (p Point) Distance(q Point) float64 {
	return p.X + p.Y + q.X + q.Y
}

// Method with a pointer receiver: Distance2 also calculates a "distance," but it operates on the original Point.
// This method works with the address of the Point, so changes to the original Point are reflected.
func (p *Point) Distance2(q Point) float64 {
	return p.X + p.Y + q.X + q.Y
}

func main() {
	// 1. Storing a simple function in a variable (method values)
	// In Go, you can assign a simple function like `Add` to a variable.
	// This is called a "function value". You can call the function through this variable.
	func1 := Add
	fmt.Printf("%T\n", func1) // Prints: func(int, int) int (the type of the Add function)
	fmt.Println(func1(2, 3))  // Output: 5 (calls the function Add(2, 3))

	// 2. Storing methods on types (method values)
	// You can store methods associated with a type in variables. The method can be called using this variable.
	// A method with a value receiver operates on a copy of the struct. A method with a pointer receiver works with the original struct.

	p := Point{1, 2} // Initialize Point p with coordinates (1, 2)
	q := Point{4, 6} // Initialize Point q with coordinates (4, 6)

	// Storing the method Distance (value receiver) in a variable.
	// This will store the method using a copy of the Point p (value receiver).
	distanceFromP := p.Distance
	// Storing the method Distance2 (pointer receiver) in a variable.
	// This stores the method which operates on the address of p (pointer receiver).
	distanceFromP2 := p.Distance2

	// Call the method using the method value.
	fmt.Println(distanceFromP(q))  // Output: 13 (uses value receiver, operating on a copy of p)
	fmt.Println(distanceFromP2(q)) // Output: 13 (uses pointer receiver, operates on original p)

	// Now let's modify p and see the difference:
	p = Point{5, 6} // Update p to a new value (5, 6)

	// `distanceFromP` uses a value receiver and holds a copy of p at the time of assignment.
	// So even after changing p, it still uses the original copy of p (1, 2).
	fmt.Println(distanceFromP(q)) // Output: 13 (doesn't update because it's a method with a value receiver)

	// `distanceFromP2` uses a pointer receiver and holds the address of p.
	// It operates on the actual instance of p, so changes to p will be reflected.
	fmt.Println(distanceFromP2(q)) // Output: 21 (reflects the updated p, which is now (5, 6))
}

// ******************** NOTE *************************
// Closure Behavior in Go:
// A closure in Go refers to a function or method that "captures" its surrounding environment.
// In most cases, closures are pass-by-reference. That is,
// the variables used by the closure will refer to the same instance of the variable in memory.

// In the Case of Method Values:

// When you assign a method with a value receiver to a variable,
// Go captures the method’s receiver as a copy.
// This means the method does not reflect changes made to the original struct after the assignment.
// In contrast, when you assign a method with a pointer receiver to a variable,
// Go captures the address (pointer) of the receiver. This means the method reflects changes made to the original struct.
