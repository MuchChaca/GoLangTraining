package main

import "fmt"

func main() {
	var name string
	fmt.Print("Your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
}
