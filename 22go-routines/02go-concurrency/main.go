package main

import "fmt"

func main() {
	go foo()
	go bar()
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
