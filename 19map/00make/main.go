package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	// var myGreeting = map[string]string{} // that works too

	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
