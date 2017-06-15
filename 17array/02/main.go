package main

import "fmt"

func main() {
	var x [58]string // if there is a number in the brackets : array else it's a slice
	// can't change the size of an array; the size is not dynamic

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(x[42])
}
