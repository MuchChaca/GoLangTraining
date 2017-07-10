package main

import "fmt"

// const : _, B, C
const (
	_ = iota      // 0
	B = iota * 10 // 1 - 10
	C = iota * 10 // 2 - 10
)

func main() {
	fmt.Println(B)
	fmt.Println(C)
}

/*
(+)iota
		- the ninth letter of the Greek alphabet
		- an extremely small amount
*/
