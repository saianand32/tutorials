package main

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

type IntSliceHolder struct {
	IntSlice
	len int
}
