package main

import "fmt"

func main() {

	if true && true {
		fmt.Println("This ran")
	}

	if true && false {
		fmt.Println("This did not run")
	}
}
