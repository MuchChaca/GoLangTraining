package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		// Just wanted to try this awesome switch system
		// could do it with if/else if/.../else too
		switch {
		case i%3 == 0 && i%5 == 0: // i%15 == 0
			fmt.Println(" -- FizzBuzz")
		case i%3 == 0:
			fmt.Println(" -- Fizz")
		case i%5 == 0:
			fmt.Println(" -- Buzz")
		default:
			fmt.Println(i)
		}
	}
}
