package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	person := Person{name: "joel", age: 23}

	fmt.Println(person)
}
