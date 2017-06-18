package main

import "fmt"

// Person : Is a regular person
type Person struct {
	Name string
	Age  int
}

// DoubleZero : struct with a stuct inside
type DoubleZero struct {
	Person
	LicenseToKill bool
}

// Greeting : Greets a person
func (p Person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

// Greeting : Greets a DoubleZero
func (dz DoubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
	p1 := Person{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := DoubleZero{
		Person: Person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()

	// embedding is real x)
}
