package main

import (
	"errors"
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
		return 0, errors.New("norgate math: square root of negative number")
	}
	// implementation
	return 42, nil
}
