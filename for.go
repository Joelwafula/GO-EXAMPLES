// this is ago program showing the use for

package main

import "fmt"

//var num int = 10

func main() {
	for i := 1; i < 24; i++ {

		if i%2 == 0 {
			continue

		}
		fmt.Println(i)

	}
}
