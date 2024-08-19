package main

import "fmt"

func main() {

	//for slices
	slice := []int{1, 2, 3, 4, 5}

	fmt.Printf("%T \n", slice)  //[]int
	fmt.Printf("%v \n", slice)  //[1 2 3 4 5]
	fmt.Printf("%#v \n", slice) //[]int{1, 2, 3, 4, 5}

	//for arrays

	arr := [3]rune{'a', 'b', 'c'}
	arr2 := [3]byte{'a', 'b', 'c'}

	fmt.Printf("%T \n", arr)  //[3]int32
	fmt.Printf("%v \n", arr)  //[97 98 99]
	fmt.Printf("%#v \n", arr) //[3]int32{97, 98, 99}

	// to print the characters
	fmt.Printf("%q \n", arr)  //['a' 'b' 'c']
	fmt.Printf("%q \n", arr2) //"abc"

}
