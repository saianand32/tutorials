package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //wg is a waitgroup and is a pointer in general.. it has 3 methods Add(), Done() and Wait()

func main() {

	websites := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://facebook.com",
	}

	for _, myUrl := range websites {
		go getStatusCode(myUrl) //can just add this but we need to wait for the goroutines to do task but main method closes by then so we need to use sync package
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(myUrl string) {
	defer wg.Done() //need to mark this wg as done once function finishes  execution

	resp, err := http.Get(myUrl)

	if err != nil {
		fmt.Println("Panic error")
	} else {
		fmt.Printf("%d statuscode for %s \n", resp.StatusCode, myUrl)
	}
}
