package main

import "fmt"

func main() {
	for i := 0; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
		// Take the string and convert it into bytes
		// there will be a list of byte for each character
		// CONVERT NOT CAST
	}
}
