package main

import "fmt"

func main() {

	var a interface{}

	a = "golang"

	//a = 2

	interfaceValue, booleanValue := a.(string)

	fmt.Println(interfaceValue)

	fmt.Println(booleanValue)
}
