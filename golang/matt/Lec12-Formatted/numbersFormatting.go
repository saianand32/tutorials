package main

import (
	"fmt"
)

func main() {
	a, b := 12, 300

	c, d := 2.3, 3.45

	fmt.Printf("%d, %d \n", a, b) // 12, 300

	//Printing hexa decimal values
	fmt.Printf("%x, %x \n", a, b)   //c, 12c
	fmt.Printf("%X, %X \n", a, b)   //C, 12C
	fmt.Printf("%#x, %#x \n", a, b) //0xc, 0x12c

	//Printing floats
	fmt.Printf("%.2f, %.1f\n", c, d) //2.30, 3.5

	//space before print
	fmt.Printf("|%6d|%6d| \n", a, b)   //|    12|   300|
	fmt.Printf("|%06d|%06d| \n", a, b) //|000012|000300|
	fmt.Printf("|%-6d|%-6d| \n", a, b) //|12    |300   |
	fmt.Printf("|%6.2f| \n", c)        //|  2.30|
}
