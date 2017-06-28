package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	a := factorial(in)
	b := factorial(in)
	c := factorial(in)
	d := factorial(in)
	e := factorial(in)
	f := factorial(in)
	g := factorial(in)
	h := factorial(in)
	i := factorial(in)
	j := factorial(in)

	for n := range merge(a, b, c, d, e, f, g, h, i, j) {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100000; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
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
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, v := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
CHALLENGE #1
-- Use the code above to execute 1,000 factorial computations and in parallel.
-- Use the "fan out / fan in" pattern to accomplish this

CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are benig used.
-- Post the percentage of your ressources being used to this discussion: https://goo.gl/BxKnOL
*/
