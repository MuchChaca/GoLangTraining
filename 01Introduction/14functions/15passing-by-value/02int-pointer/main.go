package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age) // memory address : 0xc42000e280

	changeMe(&age)

	fmt.Println(&age) // memory address : 0xc42000e280
	fmt.Println(age)  // 24 (changed)
}

func changeMe(z *int) {
	fmt.Println(z)  // memory address : 0xc42000e280
	fmt.Println(*z) // 44 (unchanged)
	*z = 24
	fmt.Println(z)  // memory address : 0xc42000e280
	fmt.Println(*z) // 24 (changed)
}
