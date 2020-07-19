package main

import "fmt"

func main() {

	//Data types
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
	//The% d template in fmt.Printf () is used to format non-decimal numeric data.

	//bollean Data Type

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	//Use% t to format the bool data using the fmt.Printf () function.

	//String Data Types

	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var number string = "32 + 34"
	fmt.Printf("number = %s", number)

}
