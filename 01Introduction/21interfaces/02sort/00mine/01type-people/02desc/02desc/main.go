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
	fmt.Println("Not Sorted\t", studyGroup)

	sort.Slice(studyGroup, func(i, j int) bool { return studyGroup[i] > studyGroup[j] })
	fmt.Println("Sorted\t\t", studyGroup)

}
