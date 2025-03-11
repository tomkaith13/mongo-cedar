package cedar_entity

import (
	"context"
	"log"
	"os"

	"github.com/cedar-policy/cedar-go"
	"github.com/tomkaith13/mongo-cedar/models"
	"github.com/tomkaith13/mongo-cedar/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func GenerateCareGiverEntity(caregiverId string) (cedar.EntityMap, error) {
	// var eMap cedar.EntityMap

	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.GetMongoClient(mongoURI)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("mydb").Collection("caregivers")

	filter := bson.M{"_id": caregiverId}
	var caregiver models.CareGiverModel
	collection.FindOne(context.TODO(), filter).Decode(&caregiver)

	logger := log.Default()
	logger.Printf("care-giver: %+v", caregiver)

	return nil, nil

}
