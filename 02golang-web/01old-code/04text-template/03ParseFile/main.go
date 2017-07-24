package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	// here we call the ParseFiles of the package level
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatal(err)
	}

	// err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// here from the type level
	tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao") // we had those to the "templates bucket"
	if err != nil {
		log.Fatal(err)
	}

	// Now we can execute anything from the templates bucket
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatal(err)
	}
}
