package main

import "fmt"

func main() {
	half := func(number int) (int, bool) {
		div := number / 2
		if number%2 == 0 {
			return div, true
		}
		return div, false
	}
	fmt.Print("half(1) returns ")
	fmt.Println(half(1))
	fmt.Print("half(2) returns ")
	fmt.Println(half(2))
}
