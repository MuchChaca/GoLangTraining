package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)

	go incrementor("Foo:\t")
	go incrementor("Bar:\t")

	wg.Wait()
	fmt.Println("Final counter:\t", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

		{ // Those are not needed but help to see the lock
			mutex.Lock() // locks the code below so no other thread can do the work below

			counter++
			fmt.Println(s, i, "\tCounter:\t", counter)

			mutex.Unlock() // then unlocks
		}
	}
	wg.Done()
}

// check if there is a race condition:
// go run -race main.go
