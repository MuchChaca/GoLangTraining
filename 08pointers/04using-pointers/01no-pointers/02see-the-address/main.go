package main

import "fmt"

func zero(x int) {
	fmt.Printf("%p\n", &x) // %p : base 16 notation, with leading 0x #godoc
	fmt.Println(&x)        // address in func zero
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}
