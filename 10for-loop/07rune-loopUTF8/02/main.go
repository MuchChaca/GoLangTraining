package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		// Same 01 but with Printf()
		fmt.Printf("%v - %v - %v\n", i, string(i), []byte(string(i)))
	}
}
