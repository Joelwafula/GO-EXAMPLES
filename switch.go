package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {

	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("three")

	}
	// this is the function to find the weekday
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("its the weekend")
	default:

		fmt.Println("its a weekday")

	}
	t := time.Now()

	switch {
	case t.Hour() < 12:

	}

}
