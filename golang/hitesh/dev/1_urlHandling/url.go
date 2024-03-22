package main

import (
	"fmt"
	"net/url"
)

const myurl = "http://localhost:3000/learn?coursename=reactjs&paymentid=hiyr66y"

func main() {
	fmt.Println("HANDLING URLS IN GOLANG")

	fmt.Println(myurl)

	//1. Parsing Url

	result, _ := url.Parse(myurl)

	// fmt.Println(result)
	fmt.Println(result.Scheme)   //http://localhost:3000/learn?coursename=reactjs&paymentid=hiyr66y
	fmt.Println(result.Host)     //localhost:3000
	fmt.Println(result.Port())   //3000
	fmt.Println(result.RawQuery) //coursename=reactjs&paymentid=hiyr66y

	queryParams := result.Query()
	fmt.Printf("%T", queryParams) //url.Values  --this means like key val pairs in map
	fmt.Println(queryParams)      //[coursename:[reactjs] paymentid:[hiyr66y]]

	fmt.Println(queryParams["paymentid"])    //[hiyr66y]
	fmt.Println(queryParams["paymentid"][0]) //hiyr66y

	for key := range queryParams {
		fmt.Printf("%v = ", key)
		fmt.Printf("%v \n", queryParams[key])
	}

	// 2. Creating url from parts

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
