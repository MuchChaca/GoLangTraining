package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// GoRoutine(1)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	// GoRoutine(2)
	go func() {
		for i := 10; i < 20; i++ {
			c <- i
		}
		wg.Done()
	}()

	// GoRoutine(3)
	go func() {
		wg.Wait()
		close(c)
	}()

	// Loop (aa*)
	for n := range c {
		fmt.Println(n)
	}
}

// the wg.Add(2) is placed better

/*
Flow execution:
-------------------------------------------------------------------------------
→ func main()
	→ create channel c
	→ create waitingGroup wg
	→ add two item to the waitingGroup wg

	→ GoRoutine(1)					→ GoRoutine(2)					→ GoRoutine(3)
		→ put a value on													→ waiting for wg
		  the chanel														  to finish
		  									→ put a value on
											  the channel					→ waiting for wg
											  									  to finish
																				→ when finished
																				  closes
	→ Loop (aa*)
		→ waits values on the channel
		  and prints it
-------------------------------------------------------------------------------
*/
