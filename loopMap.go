package main

import "fmt"

func main() {

	squaredNumber := map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25}

	for i, squared := range squaredNumber {
		fmt.Println("the pairs and keys are:", i, squared)
	}
}
