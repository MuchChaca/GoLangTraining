package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2) // means that we add 2 items tot he WaaitGroup

	go foo()
	go bar()

	wg.Wait() // we wait to the WaitGroup to get to 0
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:\t", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done() // "says done and take one away fro the WaitGroup"
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:\t", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done() // "says done and take one away fro the WaitGroup"
}
