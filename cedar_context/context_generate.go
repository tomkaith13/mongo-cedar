package cedar_context

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cedar-policy/cedar-go"
	"github.com/cedar-policy/cedar-go/types"
	"github.com/tomkaith13/mongo-cedar/models"
	"github.com/tomkaith13/mongo-cedar/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func GenerateContext(careRecepientId string, caregiverId string, capabilityId string) (*cedar.Record, error) {
	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.GetMongoClient(mongoURI)
	if err != nil {
		return nil, err
	}

	collection := client.Database("mydb").Collection("carereceipents")
	filter := bson.M{"_id": careRecepientId}
	var carereceipent models.CareReceipentModel
	collection.FindOne(context.TODO(), filter).Decode(&carereceipent)

	fmt.Printf("care receipent: %+v\n", carereceipent)

	_, ok := carereceipent.CareGiverCapabilityPermissionMap[caregiverId]
	if !ok {
		return nil, errors.New("No Caregiver found")
	}

	_, ok = carereceipent.CareGiverCapabilityPermissionMap[caregiverId][capabilityId]
	if !ok {
		return nil, errors.New("No Capability found")
	}

	// fetch permission set of a care giver
	permissionSet := carereceipent.CareGiverCapabilityPermissionMap[caregiverId][capabilityId]

	recordMap := cedar.RecordMap{}
	permissions := []types.Value{}

	for _, permission := range permissionSet.Permissions {
		permissions = append(permissions, types.String(permission))
	}

	resources := []types.Value{}
	for capId := range carereceipent.CareGiverCapabilityPermissionMap[caregiverId] {
		resources = append(resources, types.String(capId))
	}

	recordMap["impersonation"] = cedar.True
	recordMap["crId"] = types.String(careRecepientId)
	recordMap["CRCGActionSet"] = cedar.NewSet(permissions...)
	recordMap["CRCGResourceSet"] = cedar.NewSet(resources...)
	record := cedar.NewRecord(recordMap)

	return &record, nil

}
