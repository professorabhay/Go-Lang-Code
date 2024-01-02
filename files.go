package main

import (
	"fmt"
	"io"
	"os"
)

func Files() {
	fmt.Println("Hello Files in GoLang")

	content := "This need to go in file - github.com/professorabhay"

	file, err := os.Create("./github.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	defer file.Close()
	readFile("./github.txt")

}

func readFile(filename string){
	databyte , err := os.ReadFile(filename)
	checkErr(err)

	fmt.Println(string(databyte))
}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}
