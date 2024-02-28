package main

import "fmt"

func main() {
	fmt.Println("")

	num := 9
	var ptr *int = &num

	fmt.Println(ptr) //0xc00000a0e0

	num2 := 90
	ptr2 := &num2

	fmt.Println(ptr2)  //0xc00000a0e8
	fmt.Println(*ptr2) //90

	*ptr = *ptr * 2
	fmt.Println(*ptr) //18

	//simple swap function using pointers
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}
