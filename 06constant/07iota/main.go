package main

import "fmt"

const (
	_ = iota // 0
	// x << (x*x) : shift a bit
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 2 << (2 * 10)
	GB = 1 << (iota * 10) // 2 << (3 * 10)
)

func main() {
	fmt.Println("Binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n\n", KB)

	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n\n", MB)

	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}

/*
(+)iota
		- the ninth letter of the Greek alphabet
		- an extremely small amount
*/
