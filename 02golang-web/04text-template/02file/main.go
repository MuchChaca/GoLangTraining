package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close()

	// We execute the template into a file
	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", nil)
	// err = tpl.Execute(nf, nil) // or this one works too but executes only the first template
	if err != nil {
		log.Fatal(err)
	}
}
