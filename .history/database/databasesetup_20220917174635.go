package database

import(

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBset() *mongo.Client{

	client, err := mongo.NewClient(opt)

}

func UserData(client *mong.Client, collectionName string) *mong.Collection{

}

func ProductData(client *mong.Client, collectionName string) *mong.Collection{

}
