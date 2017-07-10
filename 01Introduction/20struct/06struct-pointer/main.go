package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"James", 20}

	fmt.Println("p1\t", p1)
	fmt.Printf("type\t%T\n", p1)
	fmt.Println("p1.name\t", p1.name)
	fmt.Println("p1.age\t", p1.age)
}
