package main

import (
	"fmt"
	"strings"
)

func main() {

	n1 := "programiz"
	n2 := "programiz pro"
	n3 := "programiz"

	fmt.Println(strings.Compare(n1, n2))
	fmt.Println(strings.Compare(n1, n3))

	fmt.Println(strings.Compare(n2, n3))
	fmt.Println(strings.Compare(n3, n2))

	text := "golang programming"
	sub1 := "go"
	sub2 := "golang"
	sub3 := "programming"
	sub4 := "plogramming"
	sub5 := "golang programming"

	fmt.Println(strings.Contains(text, sub1))
	fmt.Println(strings.Contains(text, sub2))
	fmt.Println(strings.Contains(text, sub3))
	fmt.Println(strings.Contains(text, sub4))
	fmt.Println(strings.Contains(text, sub5))

	text1 := "car"

	fmt.Println(strings.Replace(text1, "a", "v", 1))

	//split

	text3 := "I love programming"

	splitted := strings.Split(text3, "")

	fmt.Println(splitted)

}
