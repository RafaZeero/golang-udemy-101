package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	t := triangle{5, 10}
	s := square{2}

	fmt.Println(printArea(t))
	fmt.Println(printArea(s))
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * .5
}

func printArea(s shape) float64 {
	return s.getArea()
}
