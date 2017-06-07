package main

import "fmt"

func main() {

	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a) // almost like in C
	fmt.Printf("a's memory address - %d\n", &a) 

