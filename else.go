package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("seven is even")
	} else {
		fmt.Println("seven is odd")
	}

	if 7%2 == 0 || 8%2 == 0 {
		fmt.Println("either seven is odd or eight is even")
	}

	if n := 10; n < 0 {
		fmt.Println("n is negative")
	} else if n > 0 {
		fmt.Println("n is positive")

	} else {
		fmt.Println("the number printed is :", n)
	}

}
