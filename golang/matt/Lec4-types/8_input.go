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
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n') //this is to read the extra input of \n from fmt.scanln

	c, err := reader.ReadString('\n')

	if err == nil {
		fmt.Println(c)
	}

	d, err := reader.ReadString('\n')

	if err == nil {
		fmt.Println(d)
	}

	// to execute this with inp file  (inp.txt)
	// go run 8_input.go < inp.txt

	// *********************************************************************
	// 2. using bufio.Scanner ("Use this for all types of string ")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	x := sc.Text()

	sc.Scan()
	y := sc.Text()

	fmt.Println(x, y)
}
