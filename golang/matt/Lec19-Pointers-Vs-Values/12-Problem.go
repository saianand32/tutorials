package main

import "fmt"

func main() {
	a1 := [][2]byte{{1, 2}, {3, 4}, {4, 5}}
	a2 := [][]byte{}

	for _, v := range a1 {
		a2 = append(a2, v[:]) // v[:] converts the array to slice
	}

	fmt.Println(a1) // [[1 2] [3 4] [4 5]]
	fmt.Println(a1) // [[4 5] [4 5] [4 5]]

	//the above used to occour uptil go 1.22 as the v in for loop points to same locatiom
	// from 1.22 onward this for loop always has a copy and not pointer so it will work like this
	fmt.Println(a1) // [[1 2] [3 4] [4 5]]
	fmt.Println(a1) // [[1 2] [3 4] [4 5]]

}
