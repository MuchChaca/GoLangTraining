package main

import "fmt"

func main() {
	// Concurrency is the composition of independently executing processes,
	// while parallelism is the simulteneous execution of (possibly related) computations.
	//	// (+) Concurrency is about dealing with lots of things at once.
	//	// (+) Parallelism is about doing lots of things at once.
	foo()
	bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:\t", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:\t", i)
	}
}
