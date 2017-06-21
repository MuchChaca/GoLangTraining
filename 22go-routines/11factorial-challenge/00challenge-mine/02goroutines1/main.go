package main

import "fmt"

func main() {
	c := incrementor(10)

	f := factorial(c)

	fmt.Println(<-f)
}

func incrementor(nb int) <-chan int {
	out := make(chan int)
	go func() {
		for i := nb; i > 0; i-- {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(increment <-chan int) <-chan int {
	fact := 1
	out := make(chan int)
	go func() {
		for n := range increment {
			fact *= n
		}
		out <- fact
		close(out)
	}()
	return out
}

/*
CHALLENGE #1:
-- Use goroutines and channels to calculate factorial

CHALENGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post that answer to this discuss area: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reason of others
*/
/*
CHALENGE #2:
.
.
.
This is why, in my opinion, it is better to use goroutines and channels
*/
