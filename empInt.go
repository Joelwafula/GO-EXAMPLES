package main

import "fmt"

func showValue(i ...interface{}) {
	fmt.Println(i)
}

func main() {
	a := "mkuu"
	b := true
	c := 3

	showValue(a, b, c)

}
