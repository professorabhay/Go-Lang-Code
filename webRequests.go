package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func Requests() {
	fmt.Println("Welcome to Web-Requests in GoLang") // Simple One 

	response, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	fmt.Println(response)
	
	fmt.Printf("Response us of type %T\n", response)

	datatypes, err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(datatypes))

	defer response.Body.Close() // Request must be closed
}