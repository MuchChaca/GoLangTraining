package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body) //REadAll() returns a slice of bytes (runes)
	str := string(bs)                 // convert into string
	fmt.Println(str)
}
