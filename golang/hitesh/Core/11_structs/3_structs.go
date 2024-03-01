package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

// Defining a method for a struct
func (r rectangle) area() int {
	return r.length * r.width
}

func main() {
	rect1 := rectangle{
		length: 100,
		width:  50,
	}

	areaOfRect1 := rect1.area()

	fmt.Println(areaOfRect1)
}
