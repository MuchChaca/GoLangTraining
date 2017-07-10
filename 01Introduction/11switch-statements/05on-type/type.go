package main

import "fmt"

// switch on types
// -- normally we switch on value of variable
// -- go allows you to switch on type of variable

// Contact : create our own type
type Contact struct {
	greeting string
	name     string
}

// SwitchOnType : we'll learn more about interfaces later
func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("McLeod")
	var t = Contact{"Good to see you, ", "Todd"}
	SwitchOnType(t)
}

//_not sure if it's the complete code_\\
