package main

import "fmt"

func main() {
	// This is shorthand
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \t %T \n", a, a)
	fmt.Printf("%v \t %T \n", b, b)
	fmt.Printf("%v \t %T \n", c, c)
	fmt.Printf("%v \t %T \n", d, c)
	// %T to get the type
	// %v the value in a default format
}
