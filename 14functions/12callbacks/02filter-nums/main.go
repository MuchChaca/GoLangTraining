package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	ys := filter([]int{1, 2, 3, 4}, func(nb int) bool {
		return nb > 1
	})
	fmt.Println(ys) // [2 3 4]
}
