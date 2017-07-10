package main

import "fmt"

func main() {
	var small int
	var larger int

	fmt.Print("Enter a small number : ")
	fmt.Scan(&small)

	fmt.Print("Enter a larger number : ")
	fmt.Scan(&larger)

	fmt.Println("The remainder is ", larger%small)
}
