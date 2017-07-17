package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// here we get the Args passed with the go run command line
	// os.Args[0] is the first one an its the file to run
	name := os.Args[1] // means we can get 2 params: file and another one ?
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	str := fmt.Sprint(`
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Hello World</title>
	</head>
	<body>
		<h1>` + name + `</h1>
	</body>
</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating the file:", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
