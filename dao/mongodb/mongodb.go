package mongodb

import (
	"context"
	"fmt"
	"go-web-app/logger"
	"go-web-app/models"
	"go-web-app/settings"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

var client *mongo.Client

func Init(cfg *settings.MongodbConfig) (err error) {

	var cred options.Credential

	//cred.AuthSource = YourAuthSource
	cred.Username = cfg.Username
	cred.Password = cfg.Password

	// set client options
	clientOptions := options.Client().ApplyURI(cfg.Host).SetAuth(cred)

	// Connect to MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		zap.L().Error("init mongodb error", zap.Error(err))
		return
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		zap.L().Error("ping mongodb error", zap.Error(err))
		return
	}

	zap.L().Info("Connect to MongoDB")
	return
}

func Close() (err error) {
	err = client.Disconnect(context.TODO())

	if err != nil {
		zap.L().Error("close mongodb error", zap.Error(err))
		return
	}
	zap.L().Info("Connection to MongoDB closed.")
	return
}

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("Init settings failed, err:%v\n", err)
		return
	}
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("Init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")
	if err := Init(settings.Conf.MongodbConfig); err != nil {
		return
	}
	// ****************************
	collection := client.Database("bluebell").Collection("event")

	filter := bson.M{"event_name": "Philadelphia Eagles vs. Seattle Seahawks"}
	var result models.Event
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	if err := Close(); err != nil {
		return
	}
	return
	// ****************************
}
