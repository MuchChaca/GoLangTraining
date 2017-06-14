package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

// type prefixed by ... : i can pass in as many float64 as i want
func average(sf ...float64) float64 {
	// fmt.Println(sf)  // we get a slice (a list)
	// fmt.Printf("%T\n", sf)  // a slice of float64
	var total float64
	for _, v := range sf { // _ is the index and v the value, we don't need the index here, so we use _
		total += v
	}
	return total / float64(len(sf)) // float divided by float
}
