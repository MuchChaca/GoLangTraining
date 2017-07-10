package main

import (
	"fmt"

	"github.com/MuchChaca/GoLangTraining/04scope/01packageScope/02visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
