package main

import (
	"errors"
	"fmt"
)

func main() {

	message := "hello"

	typeErr := errors.New("There is an error present here!!")

	if message != "hello" {
		fmt.Println(typeErr)
	}

}
