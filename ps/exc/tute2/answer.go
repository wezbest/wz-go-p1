/*
Main Solution from tutorial
*/

package main

import "fmt"

// Interfaces

type shape interface {
	getArea() float64
}

// Structs

type square struct {
	sidelength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func tusol() {

	hea("Tutorial Solution")

	t := triangle{
		height: 10,
		base:   10,
	}

	s := square{
		sidelength: 10,
	}

	fmt.Println("Area of triangle:")
	printArea(t)
	fmt.Println("Area of Square:")
	printArea(s)
}
