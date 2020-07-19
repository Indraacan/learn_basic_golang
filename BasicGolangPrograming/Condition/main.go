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

	// Switch Case

	var p = 6

	switch p {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")

	}

	//Use cases for many conditions

	var pt = 9

	switch pt {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//Curly Brackets in Keyword & default cases
	var pot = 9

	switch pot {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	//Switch with if - else Style
	var num = 6

	switch {
	case num == 8:
		fmt.Println("perfect")
	case (num < 8) && (num > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// Use the Keyword # fallthrough in a # switch

	var number = 6

	switch {
	case number == 8:
		fmt.Println("perfect")
	case (number < 8) && (number > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//Nested condition selection

	var nums = 10

	if nums > 7 {
		switch nums {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if nums == 5 {
			fmt.Println("not bad")
		} else if nums == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if nums == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
