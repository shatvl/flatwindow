package routes

import (
	"flatwindow/controllers"

	"github.com/kataras/iris"
	mgo "gopkg.in/mgo.v2"
)

//DeclareRoutes for the app
func DeclareRoutes(app *iris.Application, session *mgo.Session) {
	app.PartyFunc("/auth", func(auth iris.Party) {
		auth.Post("/register", controllers.NewAuthController(session).RegisterHandler)
		auth.Post("/login", controllers.NewAuthController(session).LoginHandler)
	})
}
