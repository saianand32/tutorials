package main

import "fmt"

type school struct {
	schName string
}

type Student struct {
	Name   string
	Age    int
	Flag   bool
	Grade  complex128
	school // Embedded struct
}

func fun() (name string, age int, gender string) {
	return "sai", 21, "male"
}

func main() {

	mp := map[string]int{}

	mp["sai"]++

	fmt.Println(mp)
}
