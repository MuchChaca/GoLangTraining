package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	// bufio : a buffer for io(input/output)
	// buffer : all the data a printer needs (cache of cpu ?)
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprint(os.Stderr, "reading input:", err)
	}

	i := 0
	for k := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}

	// fmt.Println(words["test"])
	// words["test"] = "Because sometimes we need to test stuff"
	// fmt.Println(words["test"])
}
