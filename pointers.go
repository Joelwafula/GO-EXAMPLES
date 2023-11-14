package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}
func zeroPtr(iptr *int) {
	*iptr = 0
}
func main() {
	i := 1

	fmt.Println("initial is:", i)
	zeroVal(i)
	fmt.Println(i)

	zeroPtr(&i)
	fmt.Println("pointer:", &i)
	fmt.Println("zero pointer", i)

}
