package main

import "fmt"

func main() {
	c1 := incrementor("Foo:\t\t")
	c2 := incrementor("Bar:\t\t")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Counter:\t", <-c3+<-c4)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

/*
Flow
---------------------------------------------
→ incrementor(string)
	→ go func()
	→ return chan
*/

/*
(+) Race condition:
		Different goroutines trying to access and write to the same variable
(+) Deadlock:
		We have an unbuffered channel and we never have a meeting between the
		sender and the receiver
*/
