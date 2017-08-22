package main

import (
	"io"
	"net/http"
)

type snoop int

func (s snoop) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/cat/":
		io.WriteString(res, " Pussy Pussy Pussy Cat Dolls")
	case "/dog":
		io.WriteString(res, "Snoop Doggy Doggy Dogg")
	}
}

func main() {
	var h snoop
	http.ListenAndServe(":8080", h)
}
