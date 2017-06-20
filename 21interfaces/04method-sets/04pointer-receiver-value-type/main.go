package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area\t", s.area())
}

func main() {
	c := circle{5}
	// info(c) // so this one doesn't work
	info(&c)
}
