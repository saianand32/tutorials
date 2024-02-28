package main

import "fmt"

func main() {
	day := 4

	// In Go, the switch statement automatically breaks after executing a case, so you don't need to explicitly add break statements like you do in some other languages like C or Java.
	//u have to specify if u need to fall through

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
		fallthrough //no id day = 3 it prints wednesday thursday friday
	case 4:
		fmt.Println("Thursday")
		fallthrough
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	//one more way with short statement

	switch myDay := 3; myDay {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
		fallthrough //no id day = 3 it prints wednesday thursday friday
	case 4:
		fmt.Println("Thursday")
		fallthrough
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
}
