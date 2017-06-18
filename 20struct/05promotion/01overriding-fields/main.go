package main

import "fmt"

// Person : Is a regular person
type Person struct {
	First string
	Last  string
	Age   int
}

// DoubleZero : struct with a stuct inside
type DoubleZero struct {
	Person        // anonymous field (just the type here); all the inner types get promoted to the outter type
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James", // this First is overrided by the one of DoubleZero
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven", // this First overrides the First of DoubleZero
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Miss", // this First is overrided by the one of DoubleZero
			Last:  "Moneypenny",
			Age:   19,
		},
		First:         "If looks could kill", // this First overrides the First of DoubleZero
		LicenseToKill: false,
	}

	//fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(p1.First, p1.Person.First)
	fmt.Println(p2.First, p2.Person.First)
}
