package main

import (
	"fmt"
	//"strings"
)

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, nums := range nums {
		sum += nums
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 4 {
			fmt.Println("Index is : ", i)
		}

	}

	ma := map[string]int{"Joel": 1, "Wafula": 2}
	for k, v := range ma {
		fmt.Println(k, v)

	}
	for k := range ma {
		fmt.Println(k)
	}
}
