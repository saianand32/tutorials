package main

import "fmt"

func main() {

	// 1. Declare Map 3 ways

	map1 := make(map[int]int)                    //1
	var map2 map[int]int                         //2
	map3 := map[string]int{"sai": 1, "anand": 2} //3
	fmt.Println(map1, map2, map3)

	//frequency of elements
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

	// 4. making set datastructure ...go does not have set ds
	s := map[int]bool{5: true, 2: true}
	_, ok := s[6] // check for existence
	fmt.Println("ok")
	fmt.Println(ok)
	s[8] = true  // add element
	delete(s, 2) // remove element

	//5. nested maps
	mapNested := make(map[string]map[int]int)
	fmt.Println(mapNested)

	// NOTE : NOT ALL TYPES CAN BE KEYS BUT VALUES CAN BE ANY TYPE
	// only comparable tyes can be keys - boolean, numeric, string, channel, pointer and interface types also arrays and structs that contain only these types can be keys as well
	// slices, maps, functions cannot be keys as they are not comparable

}
