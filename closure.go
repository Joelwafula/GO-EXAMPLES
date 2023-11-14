package main

import "fmt"

//returning other functions without having to name them

func intSeq() func() int {
	i := 0
	return func() int {
		i++

		return i

	}
}

func main() {

	a := intSeq()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}
