package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length :number of elements referred to by the slice
	// 5 is capacity - number of elements un the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Dias!"

	// append(greeting, "suprabadham")  //=> Just this doesn't work, need to put in the var
	greeting = append(greeting, "suprabadham")

	fmt.Println(greeting[2])
}
