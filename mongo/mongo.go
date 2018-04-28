package mongo

import (
	"sync"
	"log"

	"gopkg.in/mgo.v2"
	"github.com/shatvl/flatwindow/config"
)

// Adapter is the layer for working with mongo session
type mongo struct {
	Session *mgo.Session
}

var originSession *mongo
var once sync.Once

//InitIndexes indexes in MongoDB
func InitIndexes() {
	session := Session()
	defer session.Close()

	if err := session.DB(config.Db).C("ads").EnsureIndex(mgo.Index{Key: []string{"$text:body"}}); err != nil { 
		log.Fatal(err.Error(), err) 
	}
}

// GetOriginSession creates a origin mongo session
func SetSession(s *mgo.Session) {
	once.Do(func() {
		originSession = &mongo{Session: s}
	})
}

// Session returns copy of the origin mongo session
func Session() *mgo.Session {
	return originSession.Session.Copy()
}