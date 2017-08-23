package main

import (
	"io"
	"net/http"
)

func s(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Snop Doggy Doggy Dogg")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Pussy Pussy Cat Dolls")
}

func main() {
	http.HandleFunc("/snoop", s)
	http.HandleFunc("/dolls", d)

	http.ListenAndServe(":8080", nil)
}
