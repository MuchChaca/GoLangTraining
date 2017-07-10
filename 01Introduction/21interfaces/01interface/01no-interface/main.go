package main

import "fmt"

// Square : Is a square ?
type Square struct { // we create a struct, a user defined type
	side float64
}

func (z Square) area() float64 { // then we attach a method to that type
	return z.side * z.side
}

// Triangle : #my test
type Triangle struct {
	base   float64
	height float64
}

// #my test
func (t Triangle) area() float64 {
	return 0.5 * (t.base * t.height)
}

// Shape : Is a shape?
type Shape interface { // we create a new user defined type, the underlying type is "interface"
	area() float64 // anything that has this method signature, implements the Shape interface
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	tri := Triangle{50, 50}

	info(s)
	info(tri)
}
