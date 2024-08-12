package main

import "fmt"

 
func main() {
	for i := range 10 {
		fmt.Println(i + 21)
	}
	const a = 90
	fmt.Println(a)

}
