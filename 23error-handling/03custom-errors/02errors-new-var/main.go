package main

import (
	"errors"
	"fmt"
	"log"
)

// ErrNorgateMath is an personalized error for square root on negative number
//		It is idiomatic to have "Err" at the begin of the name of the error
var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt returns the square root of the float64 in parameter
// and returns an error if the number is negative
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	// implementation
	return 42, nil
}

// see use of errors.New in standard library:
// http://golang.org/src/pkg/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go
