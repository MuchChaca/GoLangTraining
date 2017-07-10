package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	fmt.Println("Asc")
	fmt.Println("-------------------------------------")
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println("Not sorted\t", studyGroup)

	sort.Strings(studyGroup)
	fmt.Println("Sorted\t\t", studyGroup)
}
