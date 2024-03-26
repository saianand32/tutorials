package main

import (
	"encoding/json"
	"fmt"
)

// 2. Creating JSON from maps and structs using Marshal

//in the example below for structs to JSON
// In Go, field names (such as IntValue, BoolValue, etc.) and struct tags (such as json:"intValue", json:"boolValue", etc.) are both significant when encoding and decoding JSON.
// Here's why we use capitalized field names and struct tags with json::

// Capitalized Field Names:
// In Go, field names (and other identifiers) that start with a capital letter are exported, meaning they are visible and accessible outside of their own package. Fields that start with a lowercase letter are unexported, and they are only accessible within the same package.
// When encoding or decoding JSON using the encoding/json package, only exported fields will be considered. Unexported fields will be ignored. Therefore, if you want a field to be included in the JSON encoding or decoding process, you need to capitalize its name.

// json: Struct Tags:
// The json: struct tags provide instructions to the encoding/json package on how to encode and decode struct fields to and from JSON. These tags are used to specify metadata about the fields, such as the JSON key name or instructions for handling special cases.
// In your example, the json:"intValue", json:"boolValue", etc., specify the JSON key names to be used when encoding and decoding JSON. For example, the field IntValue will be encoded with the JSON key intValue, BoolValue with boolValue, and so on.

type ObjectValue struct {
	ArrayValue []int `json:"arrayValue"`
}

type Data struct {
	IntValue    int         `json:"intValue"` //here field name must be capitalized and `json:"intValue"` is called struct tags
	BoolValue   bool        `json:"boolValue"`
	StringValue string      `json:"stringValue"`
	ObjectValue ObjectValue `json:"objectValue"`
	Val         interface{} `json:"val"`
	Age         interface{} `json:"age,omitempty"` //omitempty says if Age is nil/null then omit field
	Password    string      `json:"-"`             //if u put a "-" means when u do Marshal this field is omitted and not included in json
}

func structsToJSON() {

	// 1. Structs to JSON

	data := Data{
		IntValue:    1234,
		BoolValue:   true,
		StringValue: "hello!",
		ObjectValue: ObjectValue{
			ArrayValue: []int{1, 2, 3, 4},
		},
		Val:      nil,
		Password: "abcd",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	//Note the Password is skipped as u assigned "-" in the struct tag

	fmt.Println(string(jsonData)) //{"intValue":1234,"boolValue":true,"stringValue":"hello!","objectValue":{"arrayValue":[1,2,3,4]},"val":null}
}

func mapsToJSON() {
	//Maps to JSON

	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
		"val":      nil,
		"password": "abc",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("json data: %s\n", jsonData)
}

// Slice of structs to json string
type name struct {
	Name string
	Age  int
}

func main() {
	structsToJSON()

	//can also use marshal for Slice of structs/json

	list := []name{
		{Name: "sai", Age: 22},
		{Name: "anand", Age: 33},
		{Name: "Jio", Age: 28},
	}

	jsonArray, _ := json.Marshal(list)
	fmt.Println(string(jsonArray)) //[{"Name":"sai","Age":22},{"Name":"anand","Age":33},{"Name":"Jio","Age":28}]

	//for neat print
	jsonArrayNeat, _ := json.MarshalIndent(list, "", "\t")
	fmt.Println(string(jsonArrayNeat))

}
