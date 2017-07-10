package main

import "fmt"

func main() {
	var x rune = 'a' // rune is an alias for int32
	var y int32 = 'b'

	fmt.Println(x)
	fmt.Println(y)

	// convertion: rune to string
	fmt.Println(string(x))
	fmt.Println(string(y))

}
