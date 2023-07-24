package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Hai")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // have to give these param to get the date time in this format as per the official documentation

	createdDate := time.Date(2020, time.April, 10, 23, 23, 1, 1, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
