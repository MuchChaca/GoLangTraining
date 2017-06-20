package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementor("Foo:\t")
	go incrementor("Bar:\t")
	wg.Wait()
	fmt.Println("Final counter:\t", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)

		// race:
		// counter++

		// no race:
		atomic.AddInt64(&counter, 1) // create a value that only one person can access it in the same time
		// not in the course but use this     ↓_↓
		// otherwise, go run -race main.go says it's a race condition...
		// source : http://wysocki.in/golang-concurrency-data-races/
		fmt.Println(s, i, "\tCounter:\t", atomic.LoadInt64(&counter))
	}
	wg.Done()
}

// check if there is a race condition:
// go run -race main.go
