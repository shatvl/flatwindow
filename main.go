package main

import (
	"log"
	"os"

	"github.com/kataras/iris"
	"github.com/jasonlvhit/gocron"
	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/routes"
	"github.com/shatvl/flatwindow/parsers"
	"github.com/iris-contrib/middleware/cors"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	app := iris.New()
	session, err := mgo.Dial("mongodb://shatvl:1234@ds135399.mlab.com:35399/heroku_2wq19fst")

	if err != nil {
		log.Fatal("Cannot Dial Mongo: ", err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Set origin mongo connection
	mongo.SetSession(session)
	// Declare all routes
	// Enable CORS
	app.Use(cors.New(cors.Options{
		AllowedMethods:     []string{"GET", "POST", "OPTIONS", "HEAD", "PUT", "PATCH"},
		AllowedOrigins:     []string{"*"},
		AllowedHeaders:     []string{"*"},
		ExposedHeaders:     []string{"*"},
		OptionsPassthrough: true,
		MaxAge:             3600,
	}))

	routes.DeclareRoutes(app)
	
	gocron.Every(15).Seconds().Do(parsers.NewTSParser().Parse)
	gocron.Start()
	
	port := os.Getenv("PORT")

	app.Get("/", func(ctx iris.Context) {
		ctx.Text("Works fine")
	})

	app.Run(
		iris.Addr(":" + port),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithoutVersionChecker,
		iris.WithOptimizations,
	)
}
