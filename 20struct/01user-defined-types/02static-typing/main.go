package main

import "fmt"

type foo int

func main() {
	var myAge foo = 44
	fmt.Printf("%T \t %v \n", myAge, myAge)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T \t\t %v \n", yourAge, yourAge)

	// this doesn't work :
	// fmt.Println(myAge + yourAge)  // not same type so can't add

	// this conversion works :
	// fmt.Println(int(myAge) + yourAge)
	// fmt.Println("\t\t", myAge+foo(yourAge))  // cause foo is embedded in int
}
