package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message :", msg)
	default:
		fmt.Println("there is no message to be presented")
	}

	msg := "heiiy"

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("No message sent:")
	}
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)

	case sig := <-signals:
		fmt.Println("received signals :", sig)
	default:
		fmt.Println("No activity")

	}

}
