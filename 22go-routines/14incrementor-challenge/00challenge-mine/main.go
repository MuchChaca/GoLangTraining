package main

import (
	"fmt"
	"strconv"
)

// This is actually the solution
func main() {

	c := incrementor(2)

	var count int
	for n := range c {
		count++
		fmt.Println(n)
	}
	fmt.Println("Final Count:\t", count)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprint("Process:\t"+strconv.Itoa(i)+"\tPrinting:\t", k)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	return c
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
