package main

import "fmt"

func main() {
	// For integer literals in number system

	a := 10     //decimal
	b := 0b1010 //binary
	c := 0b1010 //binary
	d := 0o12   //octal
	e := 0o12   //octal
	f := 0xA    //hexadecimal
	g := 0xA    //hexadecimal

	fmt.Println(a, b, c, d, e, f, g) //10 10 10 10 10 10 10

	// h := 0b11.0b11 //error (not allowed)

	// in complex number both real part & imag u can write in any number system unlike python where imag has to be decimal
	//in go we use only small i for imag part unlike python where we use j or J
	i := 0b11 + 0b10i

	fmt.Println(i)
}
