package main

import "fmt"

func main() {

	// if , else if & else Conditions

	var point = 8

	if point == 10 {
		fmt.Println("Perfect")
	} else if point > 5 {
		fmt.Println("Oke")
	} else if point == 4 {
		fmt.Println("Hmmmm")
	} else {
		fmt.Printf("Get Out !! %d\n", point)
	}

	// Temporary Variables if and else

	var pnt = 8840.0

	if percent := pnt / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

}
