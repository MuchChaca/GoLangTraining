package main

import (
	"fmt"
	"net/http"
)

// MyHandler a int handler
type MyHandler int

func (m MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Any code I want in this func")
}

func main() {
	var h MyHandler
	http.ListenAndServe(":8080", h)
}
