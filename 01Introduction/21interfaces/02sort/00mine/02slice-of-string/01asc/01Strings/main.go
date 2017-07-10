package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Asc")
	fmt.Println("-------------------------------------")
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Not sorted\t", s)

	sort.Strings(s)
	fmt.Println("Sorted\t\t", s)
}
