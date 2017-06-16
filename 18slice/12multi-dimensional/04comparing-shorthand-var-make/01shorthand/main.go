package main

import "fmt"

func main() {
	// shorthand
	student := []string{}
	students := [][]string{}

	student[0] = "Todd"
	// student = append(student, "Todd")  // right way

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
