package main

import "fmt"

func main() {
	// shorthand
	student := make([]string, 35)
	students := make([][]string, 35)

	student[0] = "Todd"
	// student = append(student, "Todd")  // right way

	fmt.Println(student)
	fmt.Println(students)
}
