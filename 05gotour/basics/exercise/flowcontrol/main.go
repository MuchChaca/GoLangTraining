package main

import (
	"fmt"
)

// Sqrt uses Newton's method to calculate the square root
func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		if y := z - (z*z-x)/(2*z); y == z {
			return y
		} else {
			z = y
			fmt.Println("z =", z)
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
