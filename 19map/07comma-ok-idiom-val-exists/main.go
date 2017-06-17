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

	if val, exists := myGreeting["es"]; exists {
		fmt.Println("That value exists.")
		fmt.Println("exists:", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}

	fmt.Println(myGreeting)
}
