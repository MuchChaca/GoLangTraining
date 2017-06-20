package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)

	for n := range cSum { // here we use the values on the cSum channel
		fmt.Println(n)
	}
}

func incrementor() chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)

	go func() {
		var sum int
		for n := range c { // here we use the values on the c channel
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
