package main

import (
	"encoding/json"
	"fmt"
)

// Person : A struct to turn into JSON format
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)

	// var test rune = 'A'
	// fmt.Printf("test \t %T \t %v\n", test, test)

	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}
