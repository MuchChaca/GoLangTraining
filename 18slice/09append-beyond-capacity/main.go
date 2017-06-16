package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length :number of elements referred to by the slice
	// 5 is capacity - number of elements un the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Dias!"

	greeting = append(greeting, "suprabadham")
	greeting = append(greeting, "Zao an")
	greeting = append(greeting, "Ohayou gozaimasu")
	greeting = append(greeting, "Gidday")

	fmt.Println(greeting[2])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
