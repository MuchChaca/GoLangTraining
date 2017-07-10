package main

import "fmt"

func main() {
	i := 0
	// Careful : a for loop without a condition is an infinite loop !!
	for {
		fmt.Println(i)
		i++
	}
}
