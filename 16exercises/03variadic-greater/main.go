package main

import "fmt"

func greatest(number ...int) int {
	var greatest int
	for _, i := range number {
		if i > greatest {
			greatest = i
		}
	}
	return greatest
}

func main() {
	data := []int{1, 2, 30, 5, 40}
	fmt.Println(greatest(data...))
}
