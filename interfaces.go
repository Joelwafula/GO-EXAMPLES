package main

import (
	"fmt"
	//"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rectan struct {
	width, length float64
}
type circle struct {
	radius float64
}

func (r rectan) area() float64 {
	return r.length * r.width
}
func (r rectan) perim() float64 {
	return 2*r.width + 2*r.length
}
func (c circle) areaC() float64 {
	return (c.radius * c.radius * 22) / 7
}

func main() {
	a := rectan{
		width:  20,
		length: 20,
	}
	b := circle{
		radius: 7,
	}

	c := a.area()
	d := a.perim()
	e := b.areaC()

	fmt.Println(c, d, e)

}
