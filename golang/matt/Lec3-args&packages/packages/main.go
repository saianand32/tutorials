package main

import (
	"fmt"
	morning_afternoon "greetingsProject/morningAfternoon"
	night "greetingsProject/night"
)

// note: to run this package you will have to do -----

//$ go build
//$ ./greetingsProject    ----this is nmae of you proj in go.mod which will be name of executable from $go build

func main() {

	//calling a function/ variable from same folder
	name := PersonName
	GreetPerson(name)
	fmt.Println()

	//calling function/variable from different folder/package
	morning_afternoon.GreetMorning(morning_afternoon.Person1)
	fmt.Println()

	morning_afternoon.GreetAfternoon(morning_afternoon.Person2)
	fmt.Println()

	night.GreetNight(night.Person3)
	fmt.Println()

}
