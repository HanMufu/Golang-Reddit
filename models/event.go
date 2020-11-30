package models

import (
	"github.com/globalsign/mgo/bson"
)

type Event struct {
	Id                 bson.ObjectId `bson:"_id"`
	EventId            int64         `bson:"event_id"`
	EventName          string        `bson:"event_name"`
	EventDateAndTime   int64         `bson:"event_date_and_time"`
	Cast               []string      `bson:"cast"`
	EventAddress       Address       `bson:"event_address"`
	EventType          string        `bson:"event_type"`
	TicketType         []TicketType  `bson:"ticket_type"`
	TicketingStartTime int64         `bson:"ticketing_start_time"`
	EventDescription   string        `bson:"event_description"`
}

type Address struct {
	Street  string `bson:"street"`
	City    string `bson:"city"`
	State   string `bson:"state"`
	Country string `bson:"country"`
	Zipcode string `bson:"zipcode"`
}

type TicketType struct {
	Level       string `bson:"level"`
	Price       int    `bson:"price"`
	TotalVolume int    `bson:"total_volume"`
}
