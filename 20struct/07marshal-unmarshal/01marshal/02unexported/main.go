package main

import (
	"encoding/json"
	"fmt"
)

// Person : struct to turn into JSON format
// Can't be turned into JSON cause they are not exported
type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
