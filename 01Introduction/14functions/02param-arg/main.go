package main

import "fmt"

func main() {
	greet("Jane") // arguments
	greet("John") // arguments
}

func greet(name string) { // parameter
	fmt.Println(name)
}

// greet is declared with a parameter
// when calling greet, pass in an argument
