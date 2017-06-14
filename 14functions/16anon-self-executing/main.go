package main

import "fmt"

func main() {
	func() {
		fmt.Println("I'm driving!")
	}() // the parents are here to call the func
	// oherwise it's not called
}
