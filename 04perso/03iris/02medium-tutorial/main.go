package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello World!")
	})

	app.Run(iris.Addr(":8080"))
}
