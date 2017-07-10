package main

import "fmt"

func main() {
	half := func(number int) (int, bool) {
		return number / 2, number%2 == 0
	}
	fmt.Print("half(1) returns ")
	fmt.Println(half(1))
	fmt.Print("half(2) returns ")
	fmt.Println(half(2))
}
