package main

import "fmt"

func main() {

	a := []string{"halo", "hai", "hola"}
	for i, s := range a {
		fmt.Println(i, s)
	}

	for _, b := range a {
		fmt.Println(b)
	}

	// j := 0
	intArr := [5]int{}
	for i := 0; i < len(intArr); i++ {
		fmt.Println(intArr[i])

	}

	for i := 0; i < 5; i++ {
		fmt.Println("number", i)

	}

	//Use of Keywords for Arguments Only Conditions

	var i = 0

	for i < 5 {
		fmt.Println("num", i)
		i++
	}

	var x = 0

	for x < 10 {
		fmt.Println("num", x)
		x += 2

	}

	// Use of Keyword for No Arguments

	z := 0

	for {
		fmt.Println("nums", z)

		z++
		if z == 5 {
			break
		}
	}

	// use of keyword for range
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
		fmt.Println("sum:", sum)
	}

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	//Usage keyword break and continue

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	///

	h := 1
	goto h
h:
	fmt.Println(h)

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
