package main

import "fmt"

func main() {
	name := "Todd"
	fmt.Println(name) // Todd

	changeMe(&name)

	fmt.Println(name) // Rocky
}

func changeMe(z *string) {
	*z = "Rocky"
}
