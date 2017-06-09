package main

import "fmt"

func main() {

	b := true

	// if [init a statement]; [expression to test]
	// more details on effective go
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
}
