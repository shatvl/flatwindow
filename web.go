package main

import (
	"os"

	"github.com/kataras/iris"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	app := iris.New()

	//	routes.DeclareRoutes(app, getSession())

	// Method:   GET
	app.Handle("GET", "/", func(ctx iris.Context) {
		//		parser.NewParser().Parse(getSession())
		ctx.Writef("LOL HI, AHAHAHAH")
	})

	port := os.Getenv("PORT")

	app.Run(
		iris.Addr(":"+port),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithoutVersionChecker,
		iris.WithOptimizations,
	)
}

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017/irn")

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}
