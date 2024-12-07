package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var age int
	_, err := fmt.Fscanln(os.Stdin, &name)
	_, err = fmt.Fscanln(os.Stdin, &age)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Name:", name, "Age:", age)
}
