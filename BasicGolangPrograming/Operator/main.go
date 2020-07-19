package main

import "fmt"

func main() {
	// Operator Aritmatic
	var value = (((2+6)%3)*4 - 2) / 3

	//Operator Bollean

	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	isEqual = (value != 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	isEqual = (value < 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	isEqual = (value <= 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	isEqual = (value > 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	isEqual = (value >= 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	//Logic Operator

	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
