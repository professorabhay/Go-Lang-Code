package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Switch() {
	fmt.Println("Welcome to Switch case in GoLang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")	
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		// fallthrough - It will execute all cases just beow of it 
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll the dice again")
	default:
		fmt.Println("What was that!")
	}
}