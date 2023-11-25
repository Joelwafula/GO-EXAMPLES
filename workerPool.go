// package main

// import (
// 	"fmt"
// 	"time"
// )

// func worker(id int, jobs <-chan int, results <-chan int) {
// 	for j := range jobs {

// 		fmt.Println("worker ", id, "started job", j)
// 		time.Sleep(time.Second)
// 		fmt.Println("worker", id, "finished job ", j)
// 		results <- j * 2
// 	}
// }
// func main() {
// 	const numJobs = 5

// 	jobs := make(chan int, numJobs)

// 	results := make(chan int, numJobs)

// 	for w := 1; w < 3; w++ {
// 		go worker(jobs, results, w)

// 	}
// 	for j := 1; j <= numJobs; j++ {
// 		jobs <- j

// 	}
// 	close(jobs)
// 	for a := 1; a <= numJobs; a++ {
// 		<-results
// 	}

// }
package main

import (
	"fmt"
	"time"
)

// worker function takes in an id, a jobs channel, and a results channel
func worker(id int, jobs <-chan int, results chan<- int) {
	// iterates over the jobs channel
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		// sends the result to the results channel
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	// creates a jobs channel with a buffer size of numJobs
	jobs := make(chan int, numJobs)
	// creates a results channel with a buffer size of numJobs
	results := make(chan int, numJobs)

	// creates 3 goroutines for the worker function
	for w := 1; w <= 3; w++ {

		go worker(w, jobs, results)
	}

	// sends numJobs integers to the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// close the jobs channel to signal that all jobs have been sent
	close(jobs)

	// receives numJobs integers from the results channel
	for a := 1; a <= numJobs; a++ {
		<-results
	}
	// close the results channel to signal that all results received
	close(results)
}
