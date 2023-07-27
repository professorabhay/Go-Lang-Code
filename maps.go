package main

import "fmt"

func maps() {
	fmt.Print("maps in golang\n")
	
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("js shorts for", languages["js"])

	delete(languages, "py")
	fmt.Println(languages)


	// loops are interseting in go lang [For Loop] ->
	for key, value := range languages {
		fmt.Printf("For key %s, value is %s\n", key, value)
	}

	// comma Ok sysntax for same loop ->
	for _, value := range languages{
		fmt.Printf("For key s, value is  %s\n", value)
	}



}