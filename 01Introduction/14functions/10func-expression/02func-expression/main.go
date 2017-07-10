package main

import "fmt"

func main() {

	greeting := func() { // assigning an anonymous func to a variable : func expression
		fmt.Println("Hello World!")
	}

	greeting() // only way to have a func inside a func
}
