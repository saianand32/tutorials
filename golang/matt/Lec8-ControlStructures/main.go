package main

import "fmt"

func main() {
	// Outer loop
OuterLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Outer loop:", i)
		// Inner loop
		for j := 0; j < 3; j++ {
			fmt.Println("Inner loop:", j)
			if j == 1 {
				break OuterLoop // break out of the outer loop
			} else {
				continue OuterLoop
			}
		}
	}
}
