package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=jhbj456ghb"

func URL() {
	fmt.Println("Welcome to URL's")
	fmt.Println(myurl)

	// Parsing URL
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.Hostname())
	fmt.Println(result.Query())
	fmt.Println(result.RequestURI())
	fmt.Println(result.EscapedPath())
	fmt.Println(result.Fragment)
	fmt.Println(result.User)

	qparams := result.Query()
	fmt.Printf("The type is: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	// Creating URL
	myurl2 := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "coursename=reactjs&paymentid=jhbj456ghb",
}
	fmt.Println(myurl2)
	fmt.Println(myurl2.Scheme)
	fmt.Println(myurl2.Path)
}


