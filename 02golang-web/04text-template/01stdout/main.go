package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	// say this is a template and put it in tpl
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	// err = tpl.Execute(os.Stdout, nil) // we execute the first of the templates in the terminal
	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil) // we execute the template in the terminal
	if err != nil {
		log.Fatal(err)
	}
}

/*
We can give templates any extensions we want like
[.php; .jj; .js; .mais; ...]
The custom is to give the .gohtml extension
*/
