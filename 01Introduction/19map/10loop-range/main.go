package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"en":  "Good morning.",
		"fr":  "Bonjour.",
		"es":  "Buenos dias.",
		"ita": "Bongiorno.",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
