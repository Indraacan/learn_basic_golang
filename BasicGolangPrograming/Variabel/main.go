package main

import (
	"fmt"
	"laern_basic_golang/BasicGolangPrograming/Variabel/example"
)

func main() {

	//Declaration of Variables With Data Types

	var firstName string = "john"

	var lastName string
	lastName = "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	//USe of function fmt.Printf()
	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	//Declaration of Variables without Data Types

	fName := "sony"
	lName := "nurdianto"
	fmt.Printf("halo %s %s!\n", fName, lName)

	//Multi Variable Declaration

	var first, second, third string
	first, second, third = "one", "two", "three"

	fmt.Println(first, second, third)

	uno, duo, tres := "one", "two", "three"

	fmt.Println(uno, duo, tres)

	//Variable UnderScore

	test := "i can't out"

	_ = "Learn Golang"
	_ = test
	name, _ := "Evan", "Ch"

	fmt.Println(name, test)

	//Declaration of Variables Using the ```New``` Keyword

	myname := new(string)
	fmt.Println(myname)  //will show memory adrress in the form of hexadecimal notation
	fmt.Println(*myname) // will show empty string""

	//another exampel

	fmt.Println(example.NewExample("sony", 22)) //example call from file inside example/example.go

}
