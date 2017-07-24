package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

type agent struct {
	person        // everything of person get promoted to agent
	LicenseToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := agent{
		person: person{
			Name: "James Bond",
			Age:  42,
		},
		LicenseToKill: false,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}

// check this link for more about composition:
// hhtp://www.goinggo.net/2015/09/composition-with-go.html
