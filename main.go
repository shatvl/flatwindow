package main

import (
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/routes"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/seed"
	"github.com/jasonlvhit/gocron"
	"github.com/shatvl/flatwindow/jobs"
	"github.com/shatvl/flatwindow/parsers"
	"github.com/shatvl/flatwindow/repositories"
)

func main() {
	app := iris.New()
	session, err := mgo.Dial("mongodb://bkn_admin:bkn_admin_password@185.66.68.84:27017/bkn")
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

	gocron.Every(1).Day().At("17:30").Do(parsers.NewTSParser().Parse)
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
