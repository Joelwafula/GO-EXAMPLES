package main

import "fmt"

func main() {

	var a [5]int

	fmt.Println("emp:", a)

	a[3] = 100
	fmt.Println("set:", a)

	fmt.Println("get", a[4])

	fmt.Println("len of a is:", len(a))

	b := [5]int{1, 2, 33, 44, 55}

	fmt.Println(b)

	var twoD [4][5]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = i - j
		}
	}
	fmt.Println(twoD)

}
