package main

import "fmt"

func main() {
	//representing character variables

	// use rune to rep character variables for unicode type

	var ch1 rune = 'A'
	var ch2 rune = 'S'

	fmt.Printf("character1 is %c \n", ch1) //'A'
	fmt.Printf("character2 is %c \n", ch2) //'S'

	// print Ascii val
	fmt.Printf("character1 is %d \n", ch1) //65
	fmt.Printf("character2 is %d \n", ch2) //83
}
