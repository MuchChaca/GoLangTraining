package main

import (
	"fmt"
	"net/http"
)

type myHandler int

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Snoop-Key", "this is from Snoop Dogg")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code</h1>")
}

func main() {
	var h myHandler
	http.ListenAndServe(":8080", h)
}
