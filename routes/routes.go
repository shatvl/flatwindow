package routes

import (
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/controllers"
)

//DeclareRoutes for the app
func DeclareRoutes(app *iris.Application) {
	//CORS middleware
	auth := app.Party("/auth")
	{
		auth.Post("/register", controllers.NewAuthController().RegisterHandler)
		auth.Post("/login", controllers.NewAuthController().LoginHandler)
	}

	auth.Use(cors.New(cors.Options{
		AllowedMethods:     []string{"GET", "POST", "OPTIONS", "HEAD", "PUT", "PATCH"},
		AllowedOrigins:     []string{"*"},
		AllowedHeaders:     []string{"*"},
		ExposedHeaders:     []string{"*"},
		OptionsPassthrough: true,
		MaxAge:             3600,
	}))
}
