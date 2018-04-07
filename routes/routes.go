package routes

import (
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/controllers"
	"github.com/iris-contrib/middleware/cors"
)

//DeclareRoutes for the app
func DeclareRoutes(app *iris.Application) {
	//CORS middleware
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	auth := app.Party("/auth", crs).AllowMethods(iris.MethodOptions)
	{
		auth.Post("/register", controllers.NewAuthController().RegisterHandler)
		auth.Post("/login", controllers.NewAuthController().LoginHandler)
	}
}
