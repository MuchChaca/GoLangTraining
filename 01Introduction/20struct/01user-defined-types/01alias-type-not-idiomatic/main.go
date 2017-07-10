package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Println("Type \t\t Value")
	fmt.Println("------------------------")
	fmt.Printf("%T \t %v\n", myAge, myAge)
}
