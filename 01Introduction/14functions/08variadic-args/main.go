package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	fmt.Printf("%T\n", data)
	n := average(data...) // data... gives each value of the []data
	// n := average(data)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	// func average(sf []float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
