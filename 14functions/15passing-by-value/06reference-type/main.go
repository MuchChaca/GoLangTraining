package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // []
	changeMe(m)
	fmt.Println(m) // [Todd]
	// fmt.Println(m[0])  // Todd
}

func changeMe(z []string) {
	z[0] = "Todd"
	fmt.Println(z) // [Todd]
}

// a slice is a referenced type
