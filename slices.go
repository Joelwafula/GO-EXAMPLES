package main

import (
	"fmt"
	//"slices"
)

func main() {

	var s []string
	fmt.Println("uniit", s, s == nil, len(s) == 0)

	s = make([]string, 3)

	fmt.Println("len", s, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set", s, s[1], len(s))
	//appending, adding an element to a slice

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(s)

	c := make([]string, len(s))

	copy(c, s)
	fmt.Println(c)

	//we can modify the slices in various ways

	var l = s[2:5]

	fmt.Println("s1 is:", l)

	l = s[1:5]

	fmt.Println("s2 is:", l)

	l = s[2:]
	fmt.Println("s3 is:", l)

	l = s[:5]
	fmt.Println("s4 is:", l)

	//this is a 2d array

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {

		innerLen := i + 1

		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}

	}
	fmt.Println(twoD)

	twD := make([][]int, 4)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			twD[i][j] = i + j
		}
	}
	fmt.Println(twD)

}
