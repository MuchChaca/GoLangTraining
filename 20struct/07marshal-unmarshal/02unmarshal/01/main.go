package main

import (
	"encoding/json"
	"fmt"
)

// Person : struct to unmarshal
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person

	fmt.Println("First:\t", p1.First)
	fmt.Println("Last:\t", p1.Last)
	fmt.Println("Age:\t", p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("----------------------------")
	fmt.Println("First:\t", p1.First)
	fmt.Println("Last:\t", p1.Last)
	fmt.Println("Age:\t", p1.Age)
	fmt.Printf("type:\t%T\n", p1)
}
