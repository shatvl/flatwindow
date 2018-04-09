package main

import (
	"log"

	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/config"
	//"github.com/jasonlvhit/gocron"
	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/routes"
	//"github.com/shatvl/flatwindow/parsers"
	"gopkg.in/mgo.v2"
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
	session.DB(config.Db).C("ads").EnsureIndex(mgo.Index{
		Key: []string{"rooms", "_id"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	})
	// Declare all routes
	routes.DeclareRoutes(app)
	
	//gocron.Every(15).Seconds().Do(parsers.NewTSParser().Parse)
	//gocron.Start()
	
	port := "5000" //os.Getenv("PORT")

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
