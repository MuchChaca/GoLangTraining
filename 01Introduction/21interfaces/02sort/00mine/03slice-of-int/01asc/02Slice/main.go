package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Asc")
	fmt.Println("-------------------------------------")
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println("Not sorted\t", n)

	sort.Slice(n, func(i, j int) bool { return n[i] < n[j] })
	fmt.Println("Sorted\t\t", n)
}
