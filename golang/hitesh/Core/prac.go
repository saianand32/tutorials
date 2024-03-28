package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jsonn := map[string]interface{}{}

	jsonn["sai"] = 99
	jsonn["anand"] = []int{1, 2, 3, 4}

	// newJsn, _ := json.MarshalIndent(jsonn, "", "\t")
	// fmt.Println(string(newJsn))
	min := 12
	max := 200000

	num := rand.Intn(max-min+1) + min
	fmt.Println(num)

}
