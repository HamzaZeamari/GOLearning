package main

import "fmt"

func main() {
	tr := triangle{
		base:   5,
		height: 6,
	}
	sq := square{
		sideLength: 12,
	}
	printArea(tr)
	printArea(sq)
}

type shapes interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func (tri triangle) getArea() float64 {
	return tri.base * tri.height * 0.5
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(sh shapes) {
	fmt.Println(sh.getArea())
}
