package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt returns the square root of the float64 in parameter
// and returns an error if the number is negative
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrNorgateMath := fmt.Errorf("norgate math: square root of negative number: %v", f)
		return 0, ErrNorgateMath
	}
	// implementation
	return 42, nil
}
