package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	// doesn't work, we keep the scope tight
	fmt.Println(food)
}
