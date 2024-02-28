package main

import (
	"fmt"
)

func main() {

	// 1. Declare and then add
	var slice []int    // Declaration of a slice
	fmt.Println(slice) //[]

	// 2. Declare n initialize
	slice2 := []int{1, 2, 3} // Declaration and initialization of a slice
	fmt.Println(slice2)      //[1 2 3]

	//3. append
	slice2 = append(slice2, 4, 5)
	fmt.Println(slice2) //[1 2 3 4 5]

	//4. using [:] to slice these arrays

	slice2 = []int{1, 2, 3, 4, 5}
	slice2 = append(slice2[2:])
	fmt.Println(slice2) //[3 4 5]

	slice2 = []int{1, 2, 3, 4, 5}
	slice2 = append(slice2[1:3])
	fmt.Println(slice2) //[2 3]

	// 5. Remove fields from slices

	var courses = []string{"maths", "physics", "bio", "chemistry"}
	fmt.Println(courses)

	ind := 1 //say u want to delete an element at this index

	courses = append(courses[:ind], courses[ind+1:]...)

	fmt.Println(courses) //[maths bio chemistry]

	// 6. length
	slice3 := []int{1, 2, 3, 4}
	length := len(slice3)
	fmt.Println(length) //4

	// 7. using make

	highScores := make([]int, 4)
	highScores[0] = 2
	highScores[2] = 22
	highScores[3] = 222
	// highScores[6] = 999 //coz of this line runtime err as index  out of bounds

	// but if i do
	highScores = append(highScores, 4, 5, 6, 7)
	//it auto accomodates them like an ArrayList

	fmt.Println(highScores) //[2 0 22 222 0 4 5 6 7]

	for i := 0; i < len(highScores); i++ {
		fmt.Println(highScores[i])
	}

	// 8. Another parameter in make
	highScores2 := make([]int, 5, 10)
	//slices are made on top of arrays
	//here 5 = initial length of slice
	// and 10 = capacity of the underlying array after  which fresh 10 more units need to be allocated to the slice
	fmt.Println(len(highScores2)) //5
	fmt.Println(cap(highScores2)) //10

}
