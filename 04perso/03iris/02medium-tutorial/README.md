# medium.com

> [source]("https://medium.com/go-language/iris-go-framework-mongodb-552e349eab9c")

## Install Iris
```cmd
go get -u github.com/kataras/iris
```

## Simple server
```go
package main

import (
	"github.com/kataras/iris"
)
func main() {
	app := iris.Default()
	app.Get("/", func(context iris.Context) {
		context.WriteString("Welcome to Iris Go Framework!")
	})
	app.Run(iris.Addr(":8080"))
}
```
