package main

import "fmt"

// here we are going to use maps
func main() {

	m := make(map[string]int)

	m["Joel"] = 1
	m["Wafula"] = 2
	m["simiyu"] = 3

	fmt.Println("Maps is :", m)

	m1 := m["Joel"]
	fmt.Println(m1)

	fmt.Println(len(m))

}
