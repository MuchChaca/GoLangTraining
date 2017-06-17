package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(res.Body) // bs : byte slice
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", bs) // print the byte slice as a string
}
