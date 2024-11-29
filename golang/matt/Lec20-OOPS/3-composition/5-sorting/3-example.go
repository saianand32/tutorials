package main

import (
	"fmt"
	"sort"
)

// NOTE:
// sort.Sort(data Interface)  sort.Sort takes an interface

// type Interface interface {
// 	Len() int
// 	Swap(i, j int )
// 	Less (i, j int) bool
// }

// Organ represents a body organ with a name and a weight (in grams).
type Organ struct {
	Name   string // Name of the organ
	Weight int    // Weight of the organ in grams
}

// Organs is a slice of Organ used for custom sorting.
type Organs []Organ

// Len returns the number of organs in the slice. This is required by the sort.Interface.
// The Len method tells the sort package how many elements are in the collection to sort.
func (o Organs) Len() int {
	return len(o)
}

// Swap swaps the elements with indexes i and j in the slice. This is required by the sort.Interface.
// The Swap method allows the sort package to swap the positions of two elements.
func (o Organs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

// OrgansByName is a type that wraps Organs and provides sorting by the Name field.
type OrgansByName struct {
	Organs
}

// Less compares two organs by their Name. This is required by the sort.Interface.
// It returns true if the organ at index i should come before the organ at index j based on Name.
// The Less method defines the sorting order. If it returns true, it means "i < j" for sorting purposes.
func (o OrgansByName) Less(i, j int) bool {
	return o.Organs[i].Name < o.Organs[j].Name
}

// OrgansByWeight is a type that wraps Organs and provides sorting by the Weight field.
type OrgansByWeight struct {
	Organs
}

// Less compares two organs by their Weight. This is required by the sort.Interface.
// It returns true if the organ at index i should come before the organ at index j based on Weight.
func (o OrgansByWeight) Less(i, j int) bool {
	return o.Organs[i].Weight < o.Organs[j].Weight
}

func main() {
	// Create a slice of organs with their names and weights.
	s := []Organ{
		{"brain", 1340}, // Brain, weight: 1340 grams
		{"liver", 999},  // Liver, weight: 999 grams
		{"heart", 1500}, // Heart, weight: 1500 grams
	}

	// Sort organs by name using the OrgansByName type.
	sort.Sort(OrgansByName{s})
	// Print the organs sorted by name.
	fmt.Println("Sorted by Name:", s)

	// Sort organs by weight using the OrgansByWeight type.
	sort.Sort(OrgansByWeight{s})
	// Print the organs sorted by weight.
	fmt.Println("Sorted by Weight:", s)
}

/*
** Explanation of the sort.Interface:**

In Go, the `sort` package provides a general sorting mechanism, but it requires that the collection being sorted implements the `sort.Interface`. The `sort.Interface` is an interface that defines three methods that any type can implement to be sortable by the `sort` package:

1. **Len() int**: This method returns the number of elements in the collection. It tells the sorting algorithm how many elements there are to sort.

2. **Swap(i, j int)**: This method swaps the elements at indices `i` and `j`. The sorting algorithm uses this to change the order of elements during the sorting process.

3. **Less(i, j int) bool**: This method compares two elements, the ones at indices `i` and `j`, and returns `true` if the element at `i` should come before the element at `j`. This method defines the "ordering" or "sorting" logic. If `Less(i, j)` returns `true`, the sorting algorithm considers `i` to be less than `j`, meaning `i` will appear earlier in the sorted slice.

For example:
- The `Len()` method simply returns the number of elements in the slice.
- The `Swap()` method exchanges the positions of two elements in the slice.
- The `Less()` method defines the sorting order, such as comparing organs by name (alphabetically) or by weight (numerically).

The `sort.Sort` function uses the `sort.Interface` methods to sort the collection according to the provided `Less` method. The key idea is that the `sort.Interface` gives `sort.Sort` the ability to sort **any** collection that implements this interface, making Go’s sorting mechanism very flexible and reusable.

In this example, we define two custom types: `OrgansByName` and `OrgansByWeight`, both of which implement the `sort.Interface` for sorting `Organs` by `Name` and `Weight`, respectively.
*/
