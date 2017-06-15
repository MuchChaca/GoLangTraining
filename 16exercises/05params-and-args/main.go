package main

import "fmt"

func foo(number ...int) {
	fmt.Println(number)
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
