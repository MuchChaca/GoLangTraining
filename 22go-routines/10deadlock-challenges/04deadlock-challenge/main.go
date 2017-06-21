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

	for {
		fmt.Println(<-c)
	}

	// // My solution :
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c)
	// }

}

// This prints the number, but we have received this error:
// fatal error: all goroutines are asleep - deadlock!
// Where is the deadlock?
// → We are waiting for a value at line 15, that never comes (infinity loop)

// Why are all goroutines asleep?
// → The go routine (l.8) is no longer putting values on the channel so it's sleep?

// How can we fix it?
// → Loop for only 10 times instead of infinity
