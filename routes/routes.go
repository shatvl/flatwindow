package routes

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/controllers"
	"github.com/iris-contrib/middleware/cors"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/shatvl/flatwindow/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/context"
)

//DeclareRoutes for the app
func DeclareRoutes(app *iris.Application) {
	//Enable CORS middleware
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "HEAD", "OPTIONS", "PUT"},
	})

	//Enable jwt middleware
	jwtApi := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Secret), nil
		},
		ErrorHandler: func(ctx context.Context, s string) {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(iris.Map{"error": "Unauthorized"})
		},
		SigningMethod: jwt.SigningMethodHS256,
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
		api.Post("/bid", controllers.NewBidController().BidAdHandler)
		api.Post("/me", jwtApi.Serve, controllers.NewAuthController().MeHandler)
	}

	admin := app.Party("/api/admin", crs).AllowMethods(iris.MethodOptions)
	{
		admin.Post("/products", controllers.NewAdController().GetProductsHandler)
		admin.Post("/product/feed", controllers.NewAdController().AddAdToFeed)
		admin.Post("/bids", controllers.NewBidController().GetBidsHandler)
		admin.Put("/bid", controllers.NewBidController().UpdateBidHandler)
	}

	fmt.Println(app.GetRoutes())
}
