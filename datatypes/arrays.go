package main


import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to arrays")
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Mango"

	fmt.Println("Fruits are: ", fruits)
	fmt.Println("Fruits are: ", len(fruits))

	// Single line initialization ->
	// var fruits = [4]string{"Apple", "Banana", "Mango", "Orange"}


	
}
