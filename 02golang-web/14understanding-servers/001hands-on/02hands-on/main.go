package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.side * s.side
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("The area of this shape is:", s.area())
}

func main() {
	sq := square{21}
	cr := circle{14}

	fmt.Print("Square:")
	info(sq)

	fmt.Print("Circle:")
	info(cr)
}
