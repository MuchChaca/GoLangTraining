package main

import "fmt"

func main() {

	// mySlice := []string{"a", "b", "c", "g", "m", "z"}

	// fmt.Println(mySlice)
	// // everything from 2 to 4, 4 excluded
	// fmt.Println(mySlice[2:4])  // slicing a slice
	// fmt.Println(mySlice[2])    // index access ; acessing by index
	// fmt.Println("myString"[2]) // index access ; acessing by index -> 83 -> S

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"Dias!",
		"Bongiorno",
		"Ohayo",
		"Selamat pagil!",
		"Gutten morgen",
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2])

	fmt.Print("[:2] ")
	fmt.Println(greeting[:2])

	fmt.Print("[5:] ")
	fmt.Println(greeting[5:])

	fmt.Print("[:] ")
	fmt.Println(greeting[:])
}
