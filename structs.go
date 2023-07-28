package main

import "fmt"

func structs() {
	fmt.Println("Welcome to Structs")

	// no inheritance in goLang and also No super or parent
	// but we can do composition in goLang
	Abhay := Person{"Abhay", 21, "Etawah"};
	fmt.Println(Abhay)
	fmt.Printf("%+v\n", Abhay)

}
type Person struct {
	Name string
	Age  int
	City string
}
