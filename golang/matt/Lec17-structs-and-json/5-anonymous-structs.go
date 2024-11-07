package main

import "fmt"

type person1 struct {
	Name string
	Age  int
}

type person2 struct {
	Name string
	Age  int
}

func main() {

	//1. anonymous structs s1 and s2
	s1 := struct {
		Name string
		age  int
	}{
		"sai",
		23,
	}

	s2 := struct {
		Name string
		age  int
	}{
		"anand",
		25,
	}

	//equal is possible because both are anonymous so content is directly checked and if order and content is same then will work
	s1 = s2
	fmt.Println(s1) //{anand 25}

	// 2. using named Types
	s11 := person1{
		"sai",
		22,
	}
	s12 := person2{
		"anand",
		25,
	}

	//this gives compilation error even though content of these structs are same but since they are named types
	s11 = s12

	//correct way would be (for named types)
	s11 = person1(s12)
	fmt.Println(s11) //{anand 25}
}
