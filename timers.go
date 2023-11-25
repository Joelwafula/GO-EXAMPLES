package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C

	fmt.Println("timer one is fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer two has been fired")

	}()
	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("timer two has stopped")
	}
	time.Sleep(2 * time.Second)

}
