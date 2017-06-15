package main

import "fmt"

func half(number int) (int, bool) {
	return number / 2, number%2 == 0
}

func main() {
	fmt.Print("half(1) returns ")
	fmt.Println(half(1))
	fmt.Print("half(2) returns ")
	fmt.Println(half(2))

	h, even := half(5)
	fmt.Println(h, even)
}
