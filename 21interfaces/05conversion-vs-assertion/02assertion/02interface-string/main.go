package main

import "fmt"

func main() {
	var name interface{} = "Sydney"
	// var name interface{} = 12

	// str, ok := name.(string)
	// if ok {
	// 	fmt.Printf("%T\n", str)
	// } else {
	// 	fmt.Printf("value is not a string\n")
	// }

	if str, ok := name.(string); ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
