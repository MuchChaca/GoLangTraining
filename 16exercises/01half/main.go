package main

import "fmt"

func half(number int) (int, bool) {
	div := number / 2
	if number%2 == 0 {
		return div, true
	}
	return div, false

}

func main() {
	fmt.Print("half(1) returns ")
	fmt.Println(half(1))
	fmt.Print("half(2) returns ")
	fmt.Println(half(2))
}
