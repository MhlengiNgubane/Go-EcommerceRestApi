package database

import(
	""
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBset() *mongo.Client{

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

}

func UserData(client *mong.Client, collectionName string) *mong.Collection{

}

func ProductData(client *mong.Client, collectionName string) *mong.Collection{

}
