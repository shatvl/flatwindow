package main

import (
	"log"
	"github.com/kataras/iris"
	"github.com/jasonlvhit/gocron"
	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/routes"
	"github.com/shatvl/flatwindow/parsers"
	"gopkg.in/mgo.v2"
	"github.com/shatvl/flatwindow/jobs"
	"github.com/shatvl/flatwindow/seed"
	"github.com/shatvl/flatwindow/repositories"
	"os"
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
	mongo.InitIndexes()

	//static assets
	app.Use(iris.Gzip)
	app.StaticWeb("/static/xml", "./public/xml")

	// Declare all routes
	routes.DeclareRoutes(app)

	seed.SeedAgents()

	gocron.Every(1).Day().At("19:00").Do(parsers.NewTSParser().Parse)
	gocron.Every(1).Hour().Do(jobs.NewFeed().CreateFeed, repositories.FeedTypeToName[repositories.TsType])
	gocron.Start()

	//Index route for check if build works fine
	app.Get("/", func(ctx iris.Context) {
		ctx.Text("Works fine")
	})

	port := os.Getenv("PORT")
	app.Run(
		iris.Addr(":" + port),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithoutVersionChecker,
		iris.WithOptimizations,
	)
}
