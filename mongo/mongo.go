package mongo

import (
	"sync"

	"gopkg.in/mgo.v2"
)

// Adapter is the layer for working with mongo session
type mongo struct {
	session *mgo.Session
}

var originSession *mongo
var once sync.Once

// GetOriginSession creates a origin mongo session
func SetSession(s *mgo.Session) {
	once.Do(func() {
		originSession = &mongo{session: s}
	})
}

// Session returns copy of the origin mongo session
func Session() *mgo.Session {
	return originSession.session.Copy()
}
