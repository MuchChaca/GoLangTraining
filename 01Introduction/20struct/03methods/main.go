package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// func [reciver] functionName([params]) [return type] {}
// Here we attach this function (fullName) to the type person
func (p person) fullName() string {
	return p.first + " " + p.last
}

func main() {
	p1 := person{"James", "Bond", 44}
	p2 := person{"Miss", "Moneypenny", 19}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
