package routes

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/controllers"
	"github.com/iris-contrib/middleware/cors"
)

//DeclareRoutes for the app
func DeclareRoutes(app *iris.Application) {
	//Enable CORS middleware
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	auth := app.Party("/auth", crs).AllowMethods(iris.MethodOptions)
	{
		auth.Post("/register", controllers.NewAuthController().RegisterHandler)
		auth.Post("/login", controllers.NewAuthController().LoginHandler)
	}

	api := app.Party("/api", crs).AllowMethods(iris.MethodOptions)
	{			
		api.Post("/products", controllers.NewAdController().GetProductsHandler)
		api.Get("/product/{_id:string}", controllers.NewAdController().GetProductHandler)
	}

	fmt.Println(app.GetRoutes())
}
