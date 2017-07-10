package main

import (
	"encoding/json"
	"fmt"
)

// Person : struct to turn into JSON format
type Person struct {
	First string
	Last  string `json:"-"`            // So we prevent JSON from taking this // DO NOT PUT SCPAES
	Age   int    `json:"wisdom score"` // We change Age → wisdom score        // DO NOT PUT SCPAES
}

func main() {
	p1 := Person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
