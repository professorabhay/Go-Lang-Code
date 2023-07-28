package main

import "fmt"

func functions() {
	fmt.Println("Hello")
	greeter()

	// var num1 int  = 2
	// var num2 int  = 3
	// add(num1, num2)
	// add(2,3)
	result := add(2,3)
	fmt.Println(result)

	proResult := proAdder(2,5,4,3,5,5)
	fmt.Println("Pro Result is: ", proResult)
}

func greeter(){
	fmt.Println("Hello Namaste")
}

// func add(num1 int, num2 int){
// 	fmt.Println(num1 + num2)
// }

func add(num1 int, num2 int) int { //Here int define the type of vlue we return
	return (num1 + num2)
}

func proAdder(values ...int) int {
	total := 0

	for  _, val:= range values {
		total +=val
	}
	return total
}