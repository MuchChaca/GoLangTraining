package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprint("x is greater then 10")
	}
	return fmt.Sprint("x is lower than 10")
}

// golint : if blocks ends with  a return statement, so drop this
// 			else and outdent its block
