package main

import (
	"encoding/json"
	"fmt"
)

// Person : struct to unmarshal
type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"` // What's coming as "wisdom score", we store it in "Age"
}

func main() {
	var p1 Person

	fmt.Println("First:\t", p1.First)
	fmt.Println("Last:\t", p1.Last)
	fmt.Println("Age:\t", p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("----------------------------")
	fmt.Println("First:\t\t", p1.First)
	fmt.Println("Last:\t\t", p1.Last)
	fmt.Println("wisdom score:\t", p1.Age)
	fmt.Printf("type:\t\t%T\n", p1)
}
