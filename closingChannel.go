package main

import (
	"fmt"
)

func main() {

	jobs := make(chan int, 5)

	done := make(chan bool)

	go func() {
		j, more := <-jobs
		if more {
			fmt.Println("received jobs", j)
		} else {
			fmt.Println("received all jobs")
			done <- true
			return

		}

	}()

	for j := 1; j <= 3; j++ {

		jobs <- j
		fmt.Println("sent job :", j)

	}
	close(jobs)
	fmt.Println("jobs are done:")

	<-done

}
