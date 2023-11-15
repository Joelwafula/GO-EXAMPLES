package main

import "fmt"

//initializing the function rectangle

type Rectangle func(int, int) int

type RectanglePara struct {
	length  int
	breadth int
	color   string

	rect Rectangle
}

func main() {

	result := RectanglePara{
		length:  10,
		breadth: 20,
		color:   "black",

		rect: func(length, breadth int) int {
			return length * breadth
		},
	}
	fmt.Println("color of the rectangle is: ", result.color)

	fmt.Println("area of the rectangle is:", result.rect(result.breadth, result.length))

}
