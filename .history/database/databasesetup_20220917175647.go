package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client{

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancecontext.WithTimeOut(context.Background(), 10*time.Second)
}

func UserData(client *mong.Client, collectionName string) *mong.Collection{

}

func ProductData(client *mong.Client, collectionName string) *mong.Collection{

}
