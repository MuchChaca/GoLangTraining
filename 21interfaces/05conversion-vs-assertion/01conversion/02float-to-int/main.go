package main

import "fmt"

func main() {
	var x = 12
	var y = 12.1230123

	// convertion: float64 to int
	fmt.Println(int(y) + x)
	// float to int : narrowing
}
