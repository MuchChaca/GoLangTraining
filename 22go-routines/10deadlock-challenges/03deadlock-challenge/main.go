package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// close(c)
	}()

	fmt.Println(<-c)

	// // My solution :
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c)
	// }

}

// Why does this only print zero?
// And what can you do to get it to print all 0 - 9 numbers ?

// My:
// → It's blocked at line 14 until we get a value on the channel.
//   Once we get the first value, we print it and we end the program.
//   Nothing is there to block and print the rest of the values.
