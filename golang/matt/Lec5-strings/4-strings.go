package main

import "fmt"

func main() {
	str := "€"
	fmt.Println(len(str)) //3

	str = "€uro"

	fmt.Printf("%T %[1]v\n", str)         //string €uro
	fmt.Printf("%T %[1]v\n", []rune(str)) //[]int32 [8364 117 114 111]
	fmt.Printf("%T %[1]v\n", []byte(str)) //[]uint8 [226 130 172 117 114 111]

	str = "€"
	fmt.Printf("%T %[1]v\n", []byte(str)) //[]uint8 [226 130 172]
	fmt.Println(len(str)) //3 because its the number of the bytes used to represent the string in this case 3

	// NOTE : len(str) returns physically correct answer. 
	// the logically correct answer for len("€") should be 1 but its not an ASCII number.
	// So len(str) returns logically correct len only for ascii characters.


}
