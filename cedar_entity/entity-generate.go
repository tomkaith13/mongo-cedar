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

func GenerateCareGiverEntity(caregiverId string, carereceipentId string) (cedar.EntityMap, error) {
	var eMap cedar.EntityMap

	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.GetMongoClient(mongoURI)
	if err != nil {
		return nil, err
	}

	collection := client.Database("mydb").Collection("caregivers")

	filter := bson.M{"_id": caregiverId}
	var caregiver models.CareGiverModel
	err = collection.FindOne(context.TODO(), filter).Decode(&caregiver)
	if err != nil {
		return nil, err
	}

	eMap = make(cedar.EntityMap)

	var cgEntity types.Entity
	cgEntity.UID = cedar.NewEntityUID("CareGiver", cedar.String(caregiverId))
	m := make(types.RecordMap)

	crs := []types.Value{}
	for cr := range caregiver.CareReceipentIds {
		crs = append(crs, types.String(cr))
	}
	m["cr"] = types.NewSet(crs...)

	allowedCaps := []types.Value{}
	for cap := range caregiver.AllowedResourceIds {
		allowedCaps = append(allowedCaps, types.String(cap))

	}
	m["OwnResourceSet"] = types.NewSet(allowedCaps...)

	if carereceipentId != "" {
		m["status"] = types.String(caregiver.CareReceipentInviteMap[carereceipentId])
	}

	cgEntity.Attributes = types.NewRecord(m)
	eMap[cgEntity.UID] = cgEntity

	return eMap, nil

}

func AddActionEntity(actionId string, eMap cedar.EntityMap) (cedar.EntityMap, error) {

	var actionEntity types.Entity
	actionEntity.UID = cedar.NewEntityUID("Action", cedar.String(actionId))
	m := make(types.RecordMap)
	m["name"] = types.String(actionId)
	actionEntity.Attributes = types.NewRecord(m)
	eMap[actionEntity.UID] = actionEntity

	return eMap, nil
}

func AddResourceEntity(resourceId string, eMap cedar.EntityMap) (cedar.EntityMap, error) {
	var resourceEntity types.Entity
	resourceEntity.UID = cedar.NewEntityUID("Capability", cedar.String(resourceId))
	m := make(types.RecordMap)
	m["name"] = types.String(resourceId)
	resourceEntity.Attributes = types.NewRecord(m)
	eMap[resourceEntity.UID] = resourceEntity

	return eMap, nil
}
