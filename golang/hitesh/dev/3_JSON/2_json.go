package main

import (
	"encoding/json"
	"fmt"
)

type person struct { //Remember if u want Unmarshal to work need the fields to be Capitalized(first letter capital)
	Name string
	Age  int
}

func main() {
	decodeJson()
}

func decodeJson() {

	// 1. Handling Json data from Web and convberting to struct / map using "Unmarshal"

	//we using []byte as data we recieve from net is inn this format
	jsonDataFromWeb := []byte(` 
		{
			"name": "sai",
			"age": 22
		}
	`)

	var sai person

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json Valid")
		json.Unmarshal(jsonDataFromWeb, &sai)
		fmt.Printf("%#v\n", sai)
		fmt.Println(sai.Name)
		fmt.Println(sai.Age)
	} else {
		fmt.Println("Not valid")
	}

	//some cases where u want to add data to keyval pair (no struct method using only map DS)

	var myOnlineData = make(map[string]interface{}) // when downloading from web kwy is always string but value can be anything int, bool etc so we just put an interface{} behaves like "any" in typescript
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

}
