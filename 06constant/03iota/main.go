package main

import "fmt"

const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

func main() {
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
