package main

import (
	"fmt"
	"sort"
)

func slices(){
	fmt.Println("Welcome to slices")
	var fruits = []string{"Apple", "Banana", "Mango", "Orange"}
	fmt.Println("Fruits are: ", fruits)
	fmt.Println("Fruits are: ", len(fruits))
	fmt.Println("Fruits are: ", fruits[0:2])

	// append ->
	fruits = append(fruits, "Pineapple")
	fmt.Println("Fruits are: ", fruits)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867	
	// highScore[4] = 332 // Error: OutOfBound 
	highScore = append(highScore, 332) // This will work beacuse it reallocate the size.

	fmt.Println("High Scores: ", highScore)

	sort.Ints(highScore)
	fmt.Println("Sorted High Scores: ", highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby", "golang"}
	fmt.Println("Courses: ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses: ", courses)


}
