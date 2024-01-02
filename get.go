package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	// PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http6://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count: ", byteCount)
	fmt.Println("Response Body: ", responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest(){

	requestBody := strings.NewReader(`{"name":"John Doe","occupation":"gardener"}`)

	response, err := http.Post("http://localhost:8000/post", "application/json", requestBody)

	if err != nil {
		panic(err)
	} 

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
