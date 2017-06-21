package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1 // we put something on the channel but no one receives it
	fmt.Println(<-c)

	// // My solution:
	// go func() {
	// 	c <- 1
	// }()
	// fmt.Println(<-c)
}

// This results in a deadlock.
// Can you determine why?
// And what would you do to fix it?
