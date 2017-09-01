package main

import (
	"io"
	"net/http"
)

type snoop int

func (s snoop) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/cat/":
		io.WriteString(w, " Pussy Pussy Pussy Cat Dolls")
	case "/dog":
		io.WriteString(w, "Snoop Doggy Doggy Dogg")
	}
}

func main() {
	var h snoop
	http.ListenAndServe(":8080", h)
}
