package main

import "fmt"

func loops() {
	fmt.Println("Welcome to loops in GoLang")

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// fmt.Println(days)
	
	// for d := 0; d< len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i, day := range days {
	// 	fmt.Printf("Day %d of week = %s\n", i, day)
	// }
	
	// for _, day := range days {
	// 	fmt.Printf("Day of week = %s\n", day)
	// }

	rougeValue := 1

	for rougeValue<10{

		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {
			// break;
			rougeValue++
			continue
		}
		fmt.Println(rougeValue)
		rougeValue++
	}

	lco :
		fmt.Print("hello")

}