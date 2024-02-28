package main

import "fmt"

func main() {

	arr := []string{"sai", "anand", "ankur", "utsav", "hitesh"}

	// 1. For loop
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 2. for each

	for index := range arr {
		fmt.Println(arr[index])
	}

	for index, value := range arr {
		fmt.Println(index, value)
	}

	// 3. while loop (no while loop variation of for but like while)
	i := 0
	for i < len(arr) {
		fmt.Println(arr[i])
		i++
	}

	//4. break continue and goto
	for i < 100 {
		if i == 80 {
			break
		}
		if i == 36 {
			continue
		}

		if i == 40 {
			goto lco
		}
	}

lco:
	fmt.Println("Jump to point")

}
