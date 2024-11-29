package main

import "fmt"

// Define a basic interface
type Speaker interface {
	Speak() string
}

// Another interface embedding Speaker
type Greeter interface {
	Speaker
	Greet() string
}

// demo that you can embed any type into a struct not only struct into struct
type IntSlice []int

type StringHolder struct {
	Name string
}

type IntSliceHolder struct {
	IntSlice
	StringHolder
	len int
}

func main() {
	var a IntSliceHolder
	a = IntSliceHolder{
		len:          10,
		IntSlice:     []int{1, 2, 3},
		StringHolder: StringHolder{Name: "sai"},
	}
	fmt.Printf("%#v", a)
	fmt.Printf("%#v", a.Name)
}
