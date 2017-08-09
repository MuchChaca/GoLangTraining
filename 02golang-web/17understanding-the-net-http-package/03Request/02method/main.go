package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type myHandler int

func (m myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	data := struct {
		Method      string
		Submissions url.Values // map[string][]string // can also be written like this
	}{
		req.Method,
		req.Form,
	}
	tpl.ExecuteTemplate(res, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h myHandler
	http.ListenAndServe(":8080", h)
}
