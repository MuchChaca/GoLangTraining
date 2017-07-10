package main

import "fmt"

// one way to declare it
const p string = "death & taxes"

func main() {
	// an other way to declare it
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}

// a CONSTANT is a simple unchanging value
