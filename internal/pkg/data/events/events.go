package events

import (
	"github.com/na1s/mfa/internal/pkg/config"
	"github.com/na1s/mfa/internal/pkg/data"
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	Name   string   `json:"name"`
	Chains []string `json:"chains"`
	Type   string   `json:"type"`
}

const eventCollectionName = "events"

func GetEvents(config config.Configuration) []Event {
	session, err := data.GetSession(config)
	if err != nil {
		panic(err)
	}
	defer data.Close(session)
	c := session.DB(config.GetDatabaseName()).C(eventCollectionName)

	var events []Event
	err = c.Find(bson.M{}).All(&events)
	return events
}
