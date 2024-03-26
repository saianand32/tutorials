package main

import (
	"fmt"
	"time"
)

const LoginToken = "josnfinwfw987233r" //declare a variable starting with capial letter to make it public

var MyName33 = 90 // this is allowed but walrus op not allowed
// jwt := "kokdwdj" //error cant use walrus op outside function

func main() {

	// bool

	// string

	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr

	// byte // alias for uint8

	// rune // alias for int32
	// 	 // represents a Unicode code point

	// float32 float64

	// complex64 complex128

	// ********************** 3 ways to declare a variable ************************************

	// 1. using var keyword and typing fully

	var name string = "saishwar"
	var age = 90.9 //here also go auto assigns the type
	fmt.Println(name)
	fmt.Println(age)

	//  implicit type detection (var way and no var method)

	// 2. var method (here also auto detect type)
	var myAge2 = 90
	fmt.Println(myAge2)

	// 3. No var way or using walrus operator
	// Go compiler auto detects type if you use the := (walrus operator) (no var way)
	// u cant use walrus operator outside a function u need to use the other var way

	lname := "anand"
	fmt.Println(lname)

	// Print the type of variable
	fmt.Printf("type of age : %T \n", age) //float64

	//Accessing global/outside of function variables
	fmt.Println(LoginToken)

	// declaring multiple in same line
	mil, com := 1, true
	var gg, bb = 90, "sai"
	const aa, jj = 9.99, "saish"

	// var mm int, hj bool = 3, false; //err cannot do this have to do in separate line

	fmt.Println(mil)
	fmt.Println(com)
	fmt.Println(gg)
	fmt.Println(bb)

	// *** About const ***
	// All const variables must be known at compile time else error

	const fname = "sai"
	const lasname = "anand"
	const namee = fname + lasname // this is allowed coz this is done at compile time

	const currentTime = time.Now() // error as the function will return time at runtime

	// ****** All examples of datatypes *******

	i := 90         //int
	f := 90.99      // float
	b := false      //bool
	c := 0.867 + 5i // complex128

	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(b)
	fmt.Println(c)

	var myAge int
	j := myAge //now j is also int
	fmt.Println(myAge)
	fmt.Println(j)

	//********** typecast **********
	years := 90.8
	yearsInt := int(years)
	fmt.Println(yearsInt)

	//************************************************************

	// 	int  int8  int16  int32  int64 // whole numbers

	// uint uint8 uint16 uint32 uint64 uintptr // positive whole numbers

	// float32 float64 // decimal numbers

	// complex64 complex128 // imaginary numbers (rare)\

	// 	Unless you have a good reason to, stick to the following types:
	// bool
	// string
	// int
	// uint
	// byte
	// rune
	// float64
	// complex128

	// 4. (Any type in golang)   if u want to assign a type of "any" like in typescript

	var anyVar interface{}

	anyVar = 100
	fmt.Println(anyVar) //100

	anyVar = "saishwar"
	fmt.Println(anyVar) //saishwar

}
