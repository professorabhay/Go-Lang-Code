package main

import "fmt"

func Methods() {
	fmt.Println("Hello from Methods")
	Abhay := Person{"Abhay", 21, "Etawah"};
	fmt.Println(Abhay)
	fmt.Printf("%+v\n", Abhay)
	Abhay.GetStatus()
	Abhay.NewAge()

}

type Person struct {
	Name string
	Age  int
	City string

}

func (p Person) GetStatus(){
	fmt.Println("Name Of User Is: ", p.Name)
}

func(p Person) NewAge(){
	p.Age = 30
	fmt.Println("Age of User is: ", p.Age) // it didn't change orinial value because copy of object passes
}