package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)

	// sort.StringSlice(s).Sort()  // Convert into a StringSlice then get access to the Sort()
	sort.Sort(sort.StringSlice(s)) // Apply Sort() on sort.StringSlice(s) that returns an interface
	fmt.Println(s)
}

// https://golang.org/pkg/sort/#Sort
