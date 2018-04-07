package routes

import (
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/controllers"
	"github.com/iris-contrib/middleware/cors"
)

//DeclareRoutes for the app
func DeclareRoutes(app *iris.Application) {
	//CORS middleware
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	})
	auth := app.Party("/auth", cors).AllowMethods(iris.MethodOptions) 
	{
		auth.Post("/register", controllers.NewAuthController().RegisterHandler)
		auth.Post("/login", controllers.NewAuthController().LoginHandler)
	}
}
