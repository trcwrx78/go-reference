package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{
	height float64
	base float64
}
type square struct{
	sideLength float64
}

func main() {
	tr := triangle{4, 5}
	sq := square{6}

	printArea(tr)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	at := 0.5 * t.base * t.height

	return at
}

func (s square) getArea() float64 {
	as := s.sideLength * s.sideLength

	return as
}