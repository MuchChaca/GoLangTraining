package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "foo.gohtml", req.Method)
}

func bar(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "bar.gohtml", req.URL.Path)
}

func about(res http.ResponseWriter, req *http.Request) {
	d := struct {
		FName  string
		LName  string
		RQuery string
	}{
		"Snoop",
		"Dogg",
		req.URL.RawQuery,
	}
	tpl.ExecuteTemplate(res, "about.gohtml", d)
}
