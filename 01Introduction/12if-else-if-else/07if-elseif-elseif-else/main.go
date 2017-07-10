package main

import "fmt"

func main() {
	if false {
		fmt.Println("first")
	} else if false {
		fmt.Println("second")
	} else if true {
		fmt.Println("third")
	} else {
		fmt.Println("fourth")
	}
}
