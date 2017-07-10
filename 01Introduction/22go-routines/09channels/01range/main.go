package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// Code (ab*)
	go func() {
		// Loop (aa*)
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // when a channel is closed, we can no longer put values on it; we can still receive values from it
	}() // when the channel has given all its values its done and can't give values anymore

	// Loop (bb*)
	for n := range c {
		fmt.Println(n)
	}
}

/*
It's like a relay race, the next runner waits the previous one
to pass him the baton to start running; must be syncronized
*/

/*
What happens:
---------------------------------------------------------------------------------------------
→ main starts
→ we create our channel
→ we launch the code (ab*)
		→ runs on its own go routine
			→ loops from 0 to 9
			→ puts i onto the channel
			→ as soon as i gets onto the channel it's taken by the loop (bb*)
				→ we wait for the i to be taken - should be quick
→ we are ready to receive from the unbuffered channel
		→ as soon as the loop (aa*) puts a value on c, we can use it
		→ once we use the i from the channel, its gone and an other value of i can get onto it
(+) no longer need for Sleep
---------------------------------------------------------------------------------------------
*/
