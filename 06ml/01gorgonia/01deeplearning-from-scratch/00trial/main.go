package main

import (
	"fmt"

	G "gorgonia.org/gorgonia"
)

func main() {
	// First we create a new graph
	g := G.NewGraph()

	// create a node called 'x' with the value 1
	x := G.NodeFromAny(g, 1, G.WithName("x"))

	// create a node called 'y' with the value 1
	y := G.NodeFromAny(g, 1, G.WithName("y"))

	// z = x + y
	z := G.Must(G.Add(x, y))

	// Create a VM to execute the graph
	vm := G.NewTapeMachine(g)

	// Run the VM... Errors are not checked
	vm.RunAll()

	// Print the value of z
	fmt.Printf("%v\n", z.Value())
}
