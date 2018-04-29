// file: maing.go
package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	mgo "gopkg.in/mgo.v2"

	"github.com/MuchChaca/GoLangTraining/04perso/03iris/04exp_test/controllers"
)

func main() {
	app := iris.Default()

	//* Database connection
	session, err := mgo.Dial("localhost")
	if nil != err {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Database name and collection name
	// car-db is database name car is collation name
	cGenres := session.DB("exp_go_test").C("genres")
	// c.Insert(&genres.Genre{"Classic"})

	// /
	app.Get("/", func(context iris.Context) {
		context.WriteString("Welcome to Iris Go Framework!")
	})

	// //* GET /genres
	// app.Get("/genres", func(ctx iris.Context) {
	// 	result := genres.Genre{}
	// 	err = c.Find(bson.M{}).All(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	ctx.JSON(result)
	// })

	// create a sub router and register the client-side library for the iris websockets,
	// you could skip it but iris websockets supports socket.io-like API.
	genresRouter := app.Party("/genres")

	// create our mvc application targeted to /todos relative sub path.
	genresApp := mvc.New(genresRouter)

	// // any dependencies bindings here...
	genresApp.Register(
		// todo.NewMemoryService(),
		cGenres,
	// session,
	// ws.Upgrade,
	)

	// controllers registration here...
	// genresApp.Handle(new(controllers.GenreController))
	genresApp.Register(new(controllers.GenreController))

	app.Run(iris.Addr(":8080"))
}
