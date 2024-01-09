package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("hey, maths")

	// var numA int = 2
	// var numB float64 = 4.5

	// fmt.Println("sum: ", numA + int(numB))

	// Random Number

	// fmt.Println(rand.Intn(5))

	// Random from crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

	
}