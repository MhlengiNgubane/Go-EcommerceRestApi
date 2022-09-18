package database

import (
	"context"
	"fmt"
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

	ctx, cancel := context.WithTimeOut(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb")
		return nil
	}

	fmt.Println("Successfully connected to mongodb ")
	return client

}

	var Client *mong.Client = DBSet()

func UserData(client *mong.Client, collectionName string) *mong.Collection{
	var collection *mong.Collection = client.Database("Go-Ecommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mong.Client, collectionName string) *mong.Collection{
	var productcollection *mong.Collection = client.Database("Go")
}
