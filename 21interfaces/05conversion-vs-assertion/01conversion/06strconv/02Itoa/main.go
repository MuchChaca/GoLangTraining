package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Itoa : INT to ASCII
	var x = 12
	var y = "I have this many: " + strconv.Itoa(x)

	fmt.Println(y)
	// fmt.Println("I have this many: ", strconv.Itoa(x), x)

}
