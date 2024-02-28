package main

import (
	"fmt"
	"strings"
)

// 1. Always check if a ptr is nil

func increment(name *string) {
	if name != nil { //always check for nil val
		*name = strings.ReplaceAll(*name, "a", "i")
	}
}

//2. with structs u always pass as a pointer reciever

type car struct {
	color string
}

// -- pointer reciever
func (c *car) setColor(givenColor string) {
	c.color = givenColor
}

// -- nonpointer reciever
func (c car) setColor2(givenColor string) {
	c.color = givenColor
}

func main() {
	car1 := car{color: "white"}
	car2 := car{color: "white"}

	// -- with pointerReciever
	car1.setColor("blue")
	fmt.Println(car1.color) //blue

	// -- without pointerReciever
	car2.setColor2("blue")
	fmt.Println(car2.color)
}
