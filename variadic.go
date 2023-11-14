package main

import "fmt"

func mult(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	a := mult(2, 2)
	b := mult(2, 34555, 4554)

	c := []int{2, 3, 4, 5, 6, 7, 8, 89, 9, 0, 2}
	d := mult(c...)

	fmt.Println(a, b)
	fmt.Println(d)
}
