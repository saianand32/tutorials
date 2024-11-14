package main

import "fmt"

func main() {

	outer:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break outer
			} else {
				fmt.Println(i, j)
			}
		}
	}

}
