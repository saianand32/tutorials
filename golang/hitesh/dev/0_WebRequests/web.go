package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //callers responsibility to close connection

	dataByte, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//u need to convert dataByte to String
	// Print the response body (need to co)
	// dataString := string(dataByte) // or just print like below
	fmt.Printf("Response body: %s\n", dataByte)
}
