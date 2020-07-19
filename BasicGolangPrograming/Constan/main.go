package main

import "fmt"

func main() {

	//Constan

	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	//constan use with infrance type

	const lastName = "wick"
	fmt.Print("nice to meet you ", lastName, "!\n")

	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")

}
