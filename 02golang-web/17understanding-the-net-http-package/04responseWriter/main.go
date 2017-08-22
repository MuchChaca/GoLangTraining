package main

import (
	"fmt"
	"net/http"
)

type myHandler int

func (m myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Any code you want in this func")
}

func main() {
	var h myHandler
	http.ListenAndServe(":8080", h)
}
