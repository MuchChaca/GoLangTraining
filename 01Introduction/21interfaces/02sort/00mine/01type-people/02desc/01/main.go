package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	fmt.Println("Desc")
	fmt.Println("-------------------------------------")
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Not sorted\t", studyGroup)

	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))

	fmt.Println("Sorted\t\t", studyGroup)
}
