package main

import (
	"fmt"
	"math"
)

// Square : Is a square ?
type Square struct {
	side float64
}

// Circle : another shape
type Circle struct {
	radius float64
}

// Shape : Is a shape?
type Shape interface { // we create a new user defined type, the underlying type is "interface"
	area() float64 // anything that has this method signature, implements the Shape interface
} // we read it : Square and Circle implement the Shape interface

func (s Square) area() float64 {
	return s.side * s.side
}

// which implements the Shape interface
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
	fmt.Println() // just more readable?
}

func main() {
	s := Square{10}
	c := Circle{5}

	info(s)
	info(c)
}
