package main

import "fmt"

// This func visit takes two parameters :
// - slice of int
// - a func which takes an int
func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

// // This works too ;)
// func main() {
// 	disp := func(n int) {
// 		fmt.Println(n)
// 	}
// 	visit([]int{1, 2, 3, 4}, disp)
// }

/*
(+) callback : passing a func as an argument

*/
