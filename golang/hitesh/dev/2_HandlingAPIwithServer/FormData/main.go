package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	PerformRequestFormData()
}

func PerformRequestFormData() {
	myUrl := "http://localhost:8000/postForm"

	//formdata

	formData := url.Values{}
	formData.Add("fname", "sai")
	formData.Add("lname", "anand")
	formData.Add("age", "22")

	resp, err := http.PostForm(myUrl, formData)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)        //200 OK
	fmt.Println(resp.ContentLength) //42

	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	// 1 way
	fmt.Println(string(bytes)) //{"age":"22","fname":"sai","lname":"anand"}

	// 2nd way (use of strings package)
	var respString strings.Builder
	byteCount, _ := respString.Write(bytes)

	fmt.Println(byteCount)           //42
	fmt.Println(respString.String()) //{"age":"22","fname":"sai","lname":"anand"}

}
