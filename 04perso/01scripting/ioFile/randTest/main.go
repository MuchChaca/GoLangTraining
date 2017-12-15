package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//* #2
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	{
		// Commiting changes
		defer f.Sync()
		defer fmt.Println("Defered")

		d2 := []byte{115, 111, 109, 101, 10}
		n2, err := f.Write(d2) // on the file 'f' we Write d2 | Returns the number of bytes written
		check(err)
		fmt.Printf("wrote %d bytes\n", n2)

		n3, err := f.WriteString("writes\n")
		fmt.Printf("wrote %d bytes\n", n3)

		fmt.Println("notDefered1")
	}
	fmt.Println("notDefered2")
}
