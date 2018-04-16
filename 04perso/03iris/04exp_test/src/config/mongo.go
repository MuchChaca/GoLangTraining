package config

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

// GetMongoDB returns the DB session to be used.
// CAREFUL ! It needs the following env variables to be set:
// MONGO_HOST ("localhost" for example)
// Default dbName for now is "exp_go_test"
func GetMongoDB() (*mgo.Database, error) {
	host := os.Getenv("MONGO_HOST")

	dbName := "exp_go_test"

	session, errDB := mgo.Dial(host)
	if errDB != nil {
		return nil, errDB
	}

	return session.DB(dbName), nil
}
