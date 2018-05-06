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
	"github.com/shatvl/flatwindow/controllers/admin"
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
		auth.Post("/me", jwtApi.Serve, controllers.NewAuthController().MeHandler)
	}

	api := app.Party("/api", crs).AllowMethods(iris.MethodOptions)
	{			
		api.Post("/products", controllers.NewAdController().GetProductsHandler)
		api.Get("/product/{_id:string}", controllers.NewAdController().GetProductHandler)
		api.Post("/bid", controllers.NewBidController().BidAdHandler)
	}

	//adminApi := app.Party("/api/admin", crs, middleware.AgentRoleResolverMiddleware, jwtApi.Serve).AllowMethods(iris.MethodOptions)
	//{
	//	adminApi.Post("/products", adminc.NewAdController().GetProductsHandler)
	//	adminApi.Post("/product/feed", adminc.NewAdController().AddAdToFeed)
	//	adminApi.Post("/bids", adminc.NewBidController().GetBidsHandler)
	//	adminApi.Put("/bid", adminc.NewBidController().UpdateBidHandler)
	//}

	adminAuth := app.Party("/api/admin/auth", crs).AllowMethods(iris.MethodOptions)
	{
		adminAuth.Post("/login", adminc.NewAuthController().LoginHandler)
	}

	fmt.Println(app.GetRoutes())
}
