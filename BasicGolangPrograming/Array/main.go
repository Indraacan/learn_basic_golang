package main

import "fmt"

func main() {

	//Array
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])
	fmt.Println(names)

	// Initializing Initial Array Values ​​Without Number of Elements
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("number of elemen \t:", len(numbers))

	// Array Multi Dimentsional

	var numbers1 = [2][3]int{[3]int{3, 2, 4}, [3]int{3, 4, 4}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// Looping Array Elements Using Keyword for

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	// Looping Array Elements Using Keyword for - range

	var langue = [4]string{"C", "Python", "JavaScript", "Golang"}

	for i, langue := range langue {
		fmt.Printf("elemen %d : %s\n", i, langue)
	}

	//Use of Underscore Variables _ In the for-range

	var country = [4]string{"indonesia", "singapore", "Rusia", "Korea"}

	for _, country := range country {
		fmt.Printf("name country : %s\n", country)
	}

	for i, _ := range country {
		fmt.Printf("indeks : %d\n", i)
	}
	// or
	for i := range country {
		fmt.Printf("indeks : %d\n", i)
	}
}
