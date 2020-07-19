package main

import "fmt"

type Thing struct {
	Name string
	num  int
}

func NewThing(name string, num int) *Thing {
	p := new(Thing)
	p.Name = name
	p.num = num
	return p
}

func main() {

	fmt.Println(NewThing("sony", 22))

}
