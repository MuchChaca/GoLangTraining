package vis

import "fmt"

// PrintVar is accessible outside of the package
func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName)
}
