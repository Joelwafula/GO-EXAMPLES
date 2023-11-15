package main

import "fmt"

func displayValue(i interface{}) {
	fmt.Println(i)
}

func main() {

	a := "Joel"
	b := "wafula"
	c := 32
	d := false

	displayValue(a)
	displayValue(b)
	displayValue(c)
	displayValue(d)

}
