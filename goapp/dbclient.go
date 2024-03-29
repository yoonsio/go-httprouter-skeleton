package goapp

// dbclient.go contains database-related methods

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// MongoClient holds master session and other db-related info
type MongoClient struct {
	session *mgo.Session // master session
	uri     string       // mongodb uri
	dbName  string       // database name
}

// NewMongoClient establishes connection to MongoDB database
// and returns new MongoClient object
func NewMongoClient(uri, dbName string) *MongoClient {
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return &MongoClient{session, uri, dbName}
}

// GetSession returns mgo.Session copied from
// MongoClient's master session
// Be sure to close the session after done
func (mc *MongoClient) GetSession() *mgo.Session {
	return mc.session.Copy()
}
