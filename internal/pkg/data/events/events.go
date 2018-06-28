package events

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	Name   string   `json:"name"`
	Chains []string `json:"chains"`
	Type   string   `json:"type"`
}

func Connect() []Event {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("mfa").C("event")

	var events []Event
	err = c.Find(bson.M{}).All(&events)
	return events
}
