package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func myName(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, "Ms. Marvel")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", myName)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Server error:", err)
	}
}
