package controllers

import (
	"log"

	"github.com/MuchChaca/GoLangTraining/04perso/03iris/04exp_test/genres"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// GenreController is out controller for genres
type GenreController struct {
	// Service genres.Service
	Session *mgo.Session
}

// BeforeActivation called once before the server ran, and before
// the routes and dependencies binded.
// You can bind custom things to the controller, add new methods, add middleware,
// add dependencies to the struct or the method(s) and more.
func (ctrl *GenreController) BeforeActivation(bf mvc.BeforeActivation) {
	// // this could be binded to a controller's function input argument
	// // if any, or struct field if any:
	// bf.Dependencies().Add(func(ctx iris.Context) (items []genres.Genre) {
	// 	ctx.ReadJSON(&items)
	// 	return
	// })
}

// Get handles the GET: /genres route.
func (ctrl *GenreController) Get(ctx iris.Context) genres.Genre {
	// return c.Service.Get(c.Session.ID())
	// coll := session.DB("car-db").C("car")

	// items := make([]genres.Genre, 0)

	//* GET /genres
	cGenres := ctrl.Session.DB("exp_go_test").C("genres")

	// If there are args to the request
	// if ctx.Params().Len() > 0 {
	result := genres.Genre{}
	err := cGenres.Find(bson.M{"label": ctx.Params().Get("label")}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	// items = append(items, result)
	ctx.JSON(result)
	return result
	// return result
	// } else {
	// 	// no params, we return all genres
	// 	err := cGenres.Find(bson.M{}).All(&items)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	ctx.JSON(items)
	// 	return items
	// }

	// app.Get("/genres", func(ctx iris.Context) {
	// })
}

// PostItemResponse the response data that will be returned as json
// after a post save action of all todo items.
type PostItemResponse struct {
	Success bool `json:"success"`
}

var emptyResponse = PostItemResponse{Success: false}

// // Post handles the POST: /genres route.
// func (c *GenreController) Post(newItems []genres.Genre) PostItemResponse {
// 	if err := c.Service.Save(c.Session.ID(), newItems); err != nil {
// 		return emptyResponse
// 	}

// 	return PostItemResponse{Success: true}
// }

// // GetSync --
// // TODO
// func (c *GenreController) GetSync(conn websocket.Connection) {
// 	conn.Join(c.Session.ID())
// 	conn.On("save", func() { // "save" event from client.
// 		conn.To(c.Session.ID()).Emit("saved", nil) // fire a "saved" event to the rest of the clients w.
// 	})

// 	conn.Wait()
// }
