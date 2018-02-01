package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// ExampleController serves the "/", "/ping" and "hello".
type ExampleController struct {
	iris.Controller
}

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.Controller("/", new(ExampleController))

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"))
}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *ExampleController) Get() {
	c.Ctx.HTML("<h1>Welcome</h1>")
}

// Get serves
// Methods:		GET
// Ressource: http:localhost:8080/ping
func (c *ExampleController) GetPing() {
	c.Ctx.WriteString("pong")
}

// GetHllo serves
// Method: 		GET
// Ressource: localhost"//localhost:8080/hello
func (c *ExampleController) GetHello() {
	c.Ctx.JSON(iris.Map{"message": "Hello Iris!"})
}
