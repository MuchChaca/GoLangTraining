package main

import (
	"io"
	"log"
	"net/http"
)

func myName(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Marvel")
}

func main() {
	http.HandleFunc("/marvel/", myName)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Server error:", err)
	}
}
