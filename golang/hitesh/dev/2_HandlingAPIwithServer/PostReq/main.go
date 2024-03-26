package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformPostRequest()
}

func PerformPostRequest() {
	url := "http://localhost:8000/post"

	body := strings.NewReader(`
		{
			"courseName":"golang",
			"price": 0
		}
	`)

	resp, err := http.Post(url, "application/json", body)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)        //200 OK
	fmt.Println(resp.ContentLength) //33

	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	// 1 way
	fmt.Println(string(bytes)) //{"courseName":"golang","price":0}

	// 2nd way (use of strings package)
	var respString strings.Builder
	byteCount, _ := respString.Write(bytes)

	fmt.Println(byteCount)           //33
	fmt.Println(respString.String()) //{"courseName":"golang","price":0}

}
