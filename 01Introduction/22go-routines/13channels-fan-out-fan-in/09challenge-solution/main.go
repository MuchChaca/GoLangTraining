package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// DIstribute work across multiple functions (ten goroutines) that all read from in.
	xc := fanOut(in, 10)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel

	// TROUBLESHOOTING
	fmt.Printf("%T\n", xc)
	fmt.Println("***************", len(xc))
	for _, v := range xc {
		fmt.Println("******", <-v)
	}

	var y int
	for n := range merge(xc...) {
		y++
		fmt.Println(y, "\t", n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int, n int) []<-chan int {
	var xc []<-chan int // this must be to zero value
	fmt.Println(len(xc))

	fmt.Println(xc)
	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in))
	}
	fmt.Println(xc)

	return xc
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done. This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
CHALLENGE #1
-- This code throws an error: fatal error: all goroutines are asleep - deadlock!
-- Fix  this code!
*/
