package main

import (
	"encoding/json"
	"fmt"
)

// NOTE: Private fields of struct (starting with small case are not exported so are not encoded in json)
// struct field tag `json:small` not compatible with reflect.StructTag.Get: bad syntax for struct tag value
type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
	small int      `json:small`
}

func main() {
	r := Response{Page: 1, Words: []string{"sai", "anand"}, small: 2}

	j, _ := json.Marshal(r)
	fmt.Println(string(j)) //{"page":1,"words":["sai","anand"]} here small doesnt come as its private

	r2 := Response{}

	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%+v\n", r2)
}
