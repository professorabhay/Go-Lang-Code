package main

import "fmt"

func Defer() {
	defer fmt.Println("World") // It automatically goes to the last closing bracket [last in first out]
	defer fmt.Println("one") 
	defer fmt.Println("two")
	fmt.Println("Hello")
	// defer fmt.Println("World")
	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}