package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type sage struct {
	Name  string
	Motto string
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs.",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change.",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	sages := []sage{buddha, gandhi, mlk}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
