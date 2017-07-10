package main

import "fmt"

func main() {
	x := 0
	// an anonymous function assigned to a variable (func expression)
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
Closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/

/*
(+)anonymous function
		a function without a name
		with it, we can do something like adding a function inside a function
(+)func expression
		assigning a func to a variable
*/
