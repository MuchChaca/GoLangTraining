package main

import (
	"fmt"
	"log"
)

// NorgateMathError is a error type we created
type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a nogate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := Sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

// Sqrt returns the square root of the float64 passed in and if there is, an error
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math max redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	// implementation
	return 42, nil
}

// see use of structs with error type in standard library:
// http://www.goinggp.net/2014/11/error-handling-in-go-part-ii.html
//
// http://goglang.org/pkg/net/#OpError
// http://goglang.org/src/pkg/net/dial.go
// http://goglang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
