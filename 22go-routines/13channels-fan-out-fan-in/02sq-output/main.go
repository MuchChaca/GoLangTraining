package main

import "fmt"
import "sync"

func main() {
	in := gen(2, 3)

	// FAN OUT
	// Distribute the sq work across two goroutines that both read from it.
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// Consume the merged output from multiple channels.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	fmt.Printf("TYPE OF NUMS:\t%T\n", nums) // just FYI
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	fmt.Printf("TYPE OF CS:\t%T\n", cs) // just FYI
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
FAN OUT
Multiple functions reading from the same channel until that channel is closed

FAN IN
A function can read from multiple unputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.

PATTERN
ther's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.

source:
https://blog.golang.org/pipelines
*/

/*
CHALLENGE #1
When know HOW to do fan out / fan in, but do we know WHY?
Why would we want to do fan out / fan in?

Fan out cause we have a lot of values to process and two workoers to process

*/
