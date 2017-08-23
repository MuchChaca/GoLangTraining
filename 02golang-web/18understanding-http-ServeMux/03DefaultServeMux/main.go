package main

import (
	"io"
	"net/http"
)

type snoop int

func (s snoop) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Snop Doggy Doggy Dogg")
}

type dolls int

func (d dolls) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Pussy Pussy Cat Dolls")
}

func main() {
	var s snoop
	var d dolls

	// mux := http.NewServeMux() // No longer need to create a ServeMux
	http.Handle("/snoop", s)
	http.Handle("/dolls", d)

	http.ListenAndServe(":8080", nil)
	// This means, use the default servemux
}
