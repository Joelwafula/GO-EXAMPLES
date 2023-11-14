package main

import "fmt"

// we have used the underscore to ignore the second value
func comp(a, b int) (int, int) {

	prod := a * b

	add := a + b
	return prod, add
}
func main() {
	_, b := comp(22, 22)

	fmt.Println(b)
}
