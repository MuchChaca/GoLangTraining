package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // we use make to create a channel ; unbuffered channel (no numbers in it)
	// c := make(chan int, 10) // this becomes a channel in which i can put 10 things in it ; avoid for now

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // take this value and put it on this channel ; the routines stops until something takes the value off
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) // we take off the value
			// this code here stops until it gets another value
		}
	}()

	time.Sleep(time.Second) // Sleep for a second so the code has time to run
}

/*
It's like a relay race, the next runner waits the previous one
to pass him the baton to start running; must be syncronized
*/
