package main

import (
	"fmt"
	"log"
	"os"
)

// init set up things by runing before main
func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf) // wer set all the log outputs to this file
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("err happened", err)
		log.Println("err happened", err)
		// log.Fatalln(err)
		// panic(err)
	}
}

// (+) niladic : returns nothing
