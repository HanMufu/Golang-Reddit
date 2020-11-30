package mongodb

import (
	"context"
	"go-web-app/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func GetEventById(eid int64) (event *models.Event, err error) {
	event = new(models.Event)
	filter := bson.D{{"event_id", eid}}
	//filter := bson.M{"event_id": eid}
	err = client.Database("bluebell").Collection("event").FindOne(context.TODO(), filter).Decode(&event)
	if err != nil {
		zap.L().Error("No event id matched in mongodb", zap.Error(err))
		log.Fatal(err)
		return
	}
	return
}
