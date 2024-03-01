package main

import "fmt"

// 1.  Here, func sub(x int, y int) int is known as the "function s ignature".

func sub(a int, b int) int {
	return a - b
}

// 2. Multiple parameters

func mul(a, b int) int {
	return a * b
}

func getDetails(name, gender string, age int) {
	fmt.Println(name)
	fmt.Println(gender)
	fmt.Println(age)
}

// 3. sending multiple return values

func getName(fname, lname string) (string, string, string) {
	return fname, lname, "anand"
}

// 4. can also put named return varables (below getCoords is same as getCoords2)
func getCoords() (x, y int) {
	// x and y are initialized with zero values

	return // automatically returns x and y
}

func getCoords2() (int, int) {
	var x int
	var y int
	return x, y
}

func main() {

	// for recieving multiple return values from function
	fname, lname, myName := getName("sai", "kumar")

	fmt.Println(fname)
	fmt.Println(lname)
	fmt.Println(myName)

	// if you want to ignnore some return values
	firName, _, myyName := getName("sai", "Anand")
	fmt.Println(firName)
	fmt.Println(myyName)

}
