package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %#x\n", 42, 42, 42)
	// %x: hexadecimal | %X: hex with caps | %#x: 0x + hex number
	fmt.Printf("%d - %b - %#x \t %#X \n", 42, 42, 42, 42)
}
