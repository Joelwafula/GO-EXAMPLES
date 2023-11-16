package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		message <- "ching"
	}()
	msg := message

	fmt.Println(msg)
}
