package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0xc42000e280

	var b *int = &a // Here we are referencing the memory address
	fmt.Println(b)  // 0xc42000e280
	fmt.Println(*b) // 43 // Here we are dereferencing the memory address

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
	// the * is on operator in this case
}
