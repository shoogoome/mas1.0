package utils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var MongoClient *mongo.Client
var MongoConn *mongo.Database

func InitMongoClient() {
	time.Sleep(5 * time.Second)
	err := NewMongoClient()
	if err == nil {
		fmt.Println("==> Connected to MongoDB!")
	} else {
		fmt.Println("==> Cannot connected to MongoDB! Try again after a few seconds...")
	}
	// 心跳goroutine
	go checkMongoClientConnection()
	return
}

func checkMongoClientConnection () {
	time.Sleep(3 *time.Second)
	for {
		if MongoClient == nil {
			InitMongoClient()
			return
		}
		err := MongoClient.Ping(context.Background(), nil)
		if err != nil {
			for err != nil {
				fmt.Println("==> Cannot connected to MongoDB! Try again after a few seconds...")
				time.Sleep(3 *time.Second)
				err = NewMongoClient()
			}
			fmt.Println("==> Connected to MongoDB!")
		} else {
			time.Sleep(5 *time.Second)
		}
	}
}

func NewMongoClient() (err error) {
	connectString := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		SystemConfig.MongoDB.Username,
		SystemConfig.MongoDB.Password,
		SystemConfig.MongoDB.Host,
		SystemConfig.MongoDB.Port,
	)
	mongoClientOptions := options.Client().ApplyURI(connectString)
	mongoClientOptions.SetConnectTimeout(1 * time.Second)
	mongoClientOptions.SetSocketTimeout(1 * time.Second)
	MongoClient, err = mongo.Connect(context.Background(), mongoClientOptions)

	if err != nil {
		return
	}
	err = MongoClient.Ping(context.Background(), nil)
	if err != nil {
		return
	}
	MongoConn = MongoClient.Database(SystemConfig.MongoDB.DBName)
	return
}