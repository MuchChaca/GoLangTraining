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

		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "\tCounter:\t", atomic.LoadInt64(&counter))
	}
	wg.Done()
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
