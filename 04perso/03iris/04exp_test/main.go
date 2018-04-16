package main

import (
	"os"

	"github.com/kataras/iris"
)

func main() {
	// Get running mode
	// Set to "debug" or "success"
	envMode := os.Getenv("GOMODE")

	app := iris.New()

	app.Logger().SetLevel(envMode)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Success")
	})

	app.Run(iris.Addr(":8080"))
}
