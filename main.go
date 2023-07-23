package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello world")
	
	verifyTextio()

}

func verifyTextio(){
	var username string = "abhay"
	var password string  = "5800"

	fmt.Println("Authorization: Basic", username+":"+password)

	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool

	fmt.Printf("%v\n",smsSendingLimit)
	fmt.Printf("%f\n",costPerSMS)
	fmt.Printf("%v\n",hasPermission)
	fmt.Printf("%q\n",username)

	var empty string
	empty2 := ""

	fmt.Println(empty, empty2)

	congrats := "happy birthday" // in this sorthand declaration it autometically assume the size and type as per input
	age:=20
	fmt.Println(congrats)
	fmt.Println(age) // fp is shortcut for it 
	fmt.Printf("Variable type: %T \n", congrats)

	welcome:= "Welcome"
	fmt.Println(welcome)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter your name: ")

	//Comma Ok Syntax || error ok

	// input, _ := reader.ReadString(('\n'))
	// input, err:= reader.ReadString(('\n')) something like err to store error 
	// fmt.Println("Thanks, ",input)
	// fmt.Printf("Type of input is %T ",input)

	fmt.Println("Welcome to our pizza  app")
	fmt.Println("Please rate our pizza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)
	input2, _ :=reader.ReadString('\n')

	fmt.Println("Thanks for rating, ",input2)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input2) , 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}


