package main

import "fmt"

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

}


