package main

import (
	"fmt"
)

func tt (int, string) string {
	return ""
}

func tt (string, string) string {
	return ""
}

func main() {

	a, b := 9, 10

	switch {
	case a == 1:
		fmt.Println("1")
	case b == 9:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}

	switch(a){
	case 9:
		fmt.Println("ede")
	}

}
