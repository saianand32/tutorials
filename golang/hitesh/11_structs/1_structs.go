package main

import "fmt"

// 1. Simple struct
type car struct {
	Height        int
	Weight        int
	Model, colour string
}

// 2.  Nested struct

type Wheel struct {
	Radius   int
	Material string
}

type carNested struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel
	BackWheel  Wheel
}

// 4. Anonymous defined nested struct

type plane struct {
	Wings  int
	Weight int
	// anonymous struct Wheel
	Wheel struct {
		Radius   int
		Material string
	}
}

func main() {
	car1 := car{
		Model:  "sss",
		Height: 120,
		Weight: 1000,
	}

	// 2 ways to print

	// 1.
	fmt.Println(car1) // {sss 120 1000}

	//2. better way to preint as key val pairs
	fmt.Printf("%+v", car1) //{Height:120 Weight:1000 Model:sss colour:}

	fmt.Println(car1.Model)
	fmt.Println(car1.Weight)

	car2 := carNested{}
	car2.FrontWheel.Radius = 8
	fmt.Println(car2) // {  0 0 {8 } {0 }}

	// 3. Anonymous Structs
	plane := struct {
		WingsSpan int
		Weight    int
	}{
		WingsSpan: 200,
		Weight:    1000,
	}

	fmt.Println(plane)
}
