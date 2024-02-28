package main

import "fmt"

func main() {

	// 1. Declare Arrays
	var arr [4]int

	arr[0] = 9
	arr[1] = 3
	arr[3] = 88

	fmt.Println(arr) //[9 3 0 88]

	// 2. Declare n initialize
	var arr2 = [4]int{1, 2, 3, 4}

	arr3 := [4]string{"sai", "anand", "raj", "saishwar"}

	fmt.Println(arr2)
	fmt.Println(arr3)

	// 3. Declare array without specifying size

	arr4 := [...]bool{true, false, true}
	fmt.Println(arr4)

	// 4. Length of arr

	len := len(arr)
	fmt.Println(len) //4

	// 5. Declare array of size n
	n := 10

	arr5 := make([]int, n)
	fmt.Println(arr5)

}
