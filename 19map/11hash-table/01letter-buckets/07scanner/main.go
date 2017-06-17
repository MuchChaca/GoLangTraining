package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// An artificial input source.
	const input = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean cursus cursus nunc, 
	nec eleifend augue mattis id. Proin ultrices egestas libero in lobortis. Sed sed nulla ultrices, 
	ultrices tellus in, sollicitudin tortor. Duis tristique tortor a quam vehicula, 
	accumsan efficitur sapien dignissim. Nunc sed tellus ut quam gravida dictum nec eu est. Donec eu quam viverra, 
	dapibus dolor vitae, porta nulla. Quisque eleifend interdum mauris, 
	vitae consectetur quam. Ut volutpat sem a ante egestas hendrerit. Aenean maximus lorem ac odio suscipit, 
	sit amet semper neque auctor. Donec placerat libero et neque pulvinar imperdiet.`

	scanner := bufio.NewScanner(strings.NewReader(input)) // turn a string into a Reader so we can give it to NewScanner

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords) // Each word of the scanner

	// Count the words.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
