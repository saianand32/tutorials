package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee //a  type can refer to itself as long as its with a pointer.. if u remove the pointer compilation err
	Hired  time.Time
}

func main() {

	// 1. declare struct
	var emp Employee
	emp = Employee{
		Name:   "Sai",
		Number: 90,
	}
	fmt.Printf("%#v \n", emp) // main.Employee{Name:"Sai", Number:90, Boss:(*main.Employee)(nil), Hired:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}

	//the + operator helps format the output of struct more neat
	fmt.Printf("%+v \n", emp) //{Name: Number:0 Boss:<nil> Hired:0001-01-01 00:00:00 +0000 UTC}

	//2. initialize without keys  (here u willl have to provide all the values... cannot miss any)
	emp2 := Employee{
		"sai",
		23,
		nil,
		time.Now(),
	}
	fmt.Printf("%+v \n", emp2)

	// 3. modify field
	emp.Hired = time.Now()
	fmt.Printf("%+v \n", emp)

	emp2.Boss = &emp 
}
