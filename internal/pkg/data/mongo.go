package data

import (
	"github.com/na1s/mfa/internal/pkg/config"
	mgo "gopkg.in/mgo.v2"
)

func GetSession(config config.Configuration) (*mgo.Session, error) {
	session, err := mgo.Dial(config.GetDatabaseName())
	if err != nil {
		panic(err)
	}
	return session, err
}

func Close(session *mgo.Session) {
	session.Close()
}
