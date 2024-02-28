package main

import "fmt"

func main() {

	// 1. Declare Map

	// myMap := make(map[key_datatype]value_datatype) //syntax

	myMap := make(map[string]int)

	nameSlice := []string{"sai", "anand", "sai", "ankur", "sai", "anand"}

	for i := 0; i < len(nameSlice); i++ {

		val, ok := myMap[nameSlice[i]] //this is how u need to check if a key is present in a map or not

		if ok {
			myMap[nameSlice[i]] = val + 1
		} else {
			myMap[nameSlice[i]] = 1
		}
	}
	fmt.Println(myMap)

	// 2. getting keyValue pairs

	for key, value := range myMap {
		fmt.Println(key)
		fmt.Println(value)
	}

	// 3. Delete

	delete(myMap, "sai")
	fmt.Println(myMap)
}
