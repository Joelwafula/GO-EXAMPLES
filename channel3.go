package main

import "fmt"

func main() {
	message := make(chan string, 3)

	message <- "hello"
	message <- "Wafula"
	message <- "Simiyu"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
