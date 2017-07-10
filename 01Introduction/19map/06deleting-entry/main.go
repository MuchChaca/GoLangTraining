package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"en":  "Good morning.",
		"fr":  "Bonjour.",
		"es":  "Buenos dias.",
		"ita": "Bongiorno.",
	}

	fmt.Println(myGreeting)

	delete(myGreeting, "es")
	fmt.Println(myGreeting)
}
