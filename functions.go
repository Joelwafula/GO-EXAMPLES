package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plus2(a, b, c int) int {
	return a + b + c
}

func main() {

	add := plus(1, 2)

	fmt.Println(add)

	add = plus2(32, 32, 43)

	fmt.Println(add)

}
