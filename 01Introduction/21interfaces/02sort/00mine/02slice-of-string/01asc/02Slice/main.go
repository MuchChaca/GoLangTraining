package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	fmt.Println("Asc")
	fmt.Println("-------------------------------------")

	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Not Sorted\t", s)

	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	fmt.Println("Sorted\t\t", s)

}
