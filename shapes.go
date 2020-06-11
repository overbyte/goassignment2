package main

import "fmt"

type shape interface {
	getArea() float64
	getName() string
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (triangle) getName() string {
	return "triangle"
}

func (square) getName() string {
	return "square"
}

func printArea(s shape) {
	fmt.Printf("Area of shape %v: %v\n", s.getName(), s.getArea())
}
