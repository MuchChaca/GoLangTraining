package main

import (
	"io"
	"net/http"
)

// DogHandler is a Handler
type DogHandler int

func (d DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Snopp Doggy Dogg")
}

// --------

// CatHandler is a Handler
type CatHandler int

func (c CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Pussy Pussy Cat Dolls")
}

// --------

func main() {
	var dog DogHandler
	var cat CatHandler

	mux := http.NewServeMux()
	mux.Handle("/", dog)
	mux.Handle("/cat", cat) // this must end by exactly /cat or /cat/
	// mux.Handle("/cat/", cat) // in this case even if it ends by /cat/smthing it will go through
	// run tests, it's the best way to understand the differences

	http.ListenAndServe(":8080", mux)
}

/*
if route ends in slash
	/h/ it includes anuthing beneath

if route ends in no-slash
	/h
	it only includes that
*/
