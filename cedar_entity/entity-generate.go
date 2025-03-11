package cedar_entity

import (
	"context"
	"os"

	"github.com/cedar-policy/cedar-go"
	"github.com/cedar-policy/cedar-go/types"
	"github.com/tomkaith13/mongo-cedar/models"
	"github.com/tomkaith13/mongo-cedar/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func GenerateCareGiverEntity(caregiverId string) (cedar.EntityMap, error) {
	var eMap cedar.EntityMap

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

	eMap = make(cedar.EntityMap)

	var cgEntity types.Entity
	cgEntity.UID = cedar.NewEntityUID("CareGiver", cedar.String(caregiverId))
	m := make(types.RecordMap)
	// m["firstName"] = types.String(caregiver.FirstName)
	// m["lastName"] = types.String(caregiver.LastName)
	// m["email"] = types.String(caregiver.Email)
	// m["phone"] = types.String(caregiver.Phone)

	crs := []types.Value{}

	for cr := range caregiver.CareReceipentIds {
		crs = append(crs, types.String(cr))
	}

	m["cr"] = types.NewSet(crs...)

	cgEntity.Attributes = types.NewRecord(m)
	eMap[cgEntity.UID] = cgEntity

	return eMap, nil

}
