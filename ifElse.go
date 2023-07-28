package main

import "fmt"

func ifElse() {
	fmt.Println("Welcome to If-Else")

	loginCount := 23

	var result string

	if loginCount < 23 {
		//{	 Can'nt write like this, it give error

		result = "Regular"
	} else if loginCount > 23 {
		result = "Experienced"
	} else {
		result = "Beginner"
	}

	fmt.Println("Result is: ", result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num<10{
		fmt.Println("Number is less than 10")
	}else{
		fmt.Println("Number is not less than 10")
	}

	// if err != nil {
	// 	fmt.Println("Error Occured")
	// }

}
