package main

import "fmt"

func main() {
	name := "Yo mama"

	fmt.Println(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Hello World!</title>
</head>
<body>
	<h1>` + name + `</h1>
</body>
</html>
	`)

	// run it with:
	// go run main.go > index.html
	// so we print all of it in this file index.html
	// need to refresh to see them sometimes
	// you can use ls and cat in the terminal
}
