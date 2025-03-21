package mongo

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func GetMongoClient(mongoURI string) (*mongo.Client, error) {
	if MongoClient != nil {
		return MongoClient, nil
	}
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	passwd := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	clientOptions := options.Client().ApplyURI(mongoURI).SetAuth(options.Credential{
		Username:   user,
		Password:   passwd,
		AuthSource: "admin",
	})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	MongoClient = client
	return client, err
}
