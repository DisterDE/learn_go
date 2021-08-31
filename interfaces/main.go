package main

import "fmt"

func main() {
	shapes := []shape{
		square{sideLength: 5.0},
		triangle{base: 5, height: 5},
	}

	for _, s := range shapes {
		printArea(s)
	}
}

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
