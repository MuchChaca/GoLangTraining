package main

import "fmt"

// const : A, B, C
const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

func main() {
	// D : a const = iota
	const D = iota

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
}

/*
(+)iota
		- the ninth letter of the Greek alphabet
		- an extremely small amount
*/
