// file: maing.go
package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"

	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()

	// configure the http sessions.
	sess := sessions.New(sessions.Config{
		Cookie: "exp_session",
	})

	// // configure the websocket server.
	// ws := websocket.New(websocket.Config{})

	// create a sub router and register the client-side library for the iris websockets,
	// you could skip it but iris websockets supports socket.io-like API.
	// booksRouter := app.Party("/books")
	genresRouter := app.Party("/genres")

	// create our mvc application targeted to /books relative sub path.
	// booksApp := mvc.New(booksRouter)
	genresApp := mvc.New(genresRouter)

	// any dependencies bindings here...
	// booksApp.Register(
	// 	book.NewMemoryService(),
	// 	sess.Start,
	// 	ws.Upgrade,
	// )
	genresApp.Register(
		genre.NewMemoryService(),
		sess.Start,
		ws.Upgrade,
	)

	// controllers registration here...
	// booksApp.Handle(new(controllers.BookController))
	genresApp.Handle(new(controllers.GenreController))

	// start the web server at http://localhost:8080
	app.Run(iris.Addr(":8080"), iris.WithoutVersionChecker)
}
