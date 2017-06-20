package main

import "fmt"

func main() {

	var name interface{} = 12

	if str, ok := name.(string); ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
