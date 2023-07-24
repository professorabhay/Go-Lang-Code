package main

import "fmt"

func main() {
	fmt.Println("Welcome to pinters")
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	b := &a // b is a pointer variable that points at the memory location of 'a
	fmt.Println(b)
	fmt.Println(*b) // *b is the value stored at the memory location b
	*b = 43
	fmt.Println(a)
	fmt.Println(*b)

}