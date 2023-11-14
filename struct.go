package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{
		name: "Joel",
		age:  20,
	}
	return &p
}

func main() {

	fmt.Println(person{"joel", 3})

	s := newPerson("")

	fmt.Println(s)

}
