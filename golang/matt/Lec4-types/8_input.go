package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a int
	var b bool
	var c string

	//for any input other than space separated string use fmt.Scanln

	fmt.Scanln(&a, &b)
	// fmt.Scanln(&a)
	// fmt.Scanln(&b)

	// use bufio for string with spaces
	sc := bufio.NewReader(os.Stdin)
	sc.ReadString('\n') //this is to read the extra input of \n from fmt.scanln

	c, err := sc.ReadString('\n')

	if err == nil {
		fmt.Println(c)
	}

	d, err := sc.ReadString('\n')

	if err == nil {
		fmt.Println(d)
	}

	// to execute this with inp file  (inp.txt)
	// go run 8_input.go < inp.txt
}
