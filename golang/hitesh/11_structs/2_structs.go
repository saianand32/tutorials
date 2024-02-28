package main

import "fmt"

// Embedded structs
type car struct {
	Make  string
	Model string
}

type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	car
	BedSize int
}

type Student struct {
	Name  string
	Age   int
	Flag  bool
	Grade complex128
}

func main() {
	truck1 := truck{}
	truck1.Model = "modv123" // directly acces the car properties as top level thats how its different from nested structs

	//Default values in structs
	stud1 := Student{}
	fmt.Println(stud1.Name)  //""
	fmt.Println(stud1.Age)   //0
	fmt.Println(stud1.Flag)  //false
	fmt.Println(stud1.Grade) //0+0i
}
