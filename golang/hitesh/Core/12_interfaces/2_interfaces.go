package main

// 1. A struct can implement multiple interafaces
// it just needs to have alll methods of each interface

// 2. parametrized interface functions

type interface1 interface {
	fun(string, string) int
}

//named params
type interface2 interface {
	fun(name string, gender string) (age int)
}

func main() {

}
