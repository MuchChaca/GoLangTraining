package main

import "fmt"

// Here we are returning a func
func wrapper() func() int {
	x := 0
	// here the return + the func we are returning
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
