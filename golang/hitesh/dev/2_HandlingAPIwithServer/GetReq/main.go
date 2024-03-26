package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
}

func PerformGetRequest() {
	url := "http://localhost:8000/get"

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)        //200 OK
	fmt.Println(resp.ContentLength) //43

	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	// 1 way
	fmt.Println(string(bytes)) //{"message":"Hello from learnCodeonline.in"}

	// 2nd way (use of strings package)
	var respString strings.Builder
	byteCount, _ := respString.Write(bytes)

	fmt.Println(byteCount)           //43
	fmt.Println(respString.String()) //{"message":"Hello from learnCodeonline.in"}

}
