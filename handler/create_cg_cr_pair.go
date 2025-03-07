package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/tomkaith13/mongo-cedar/models"
	"github.com/tomkaith13/mongo-cedar/mongo"
)

func CreateCareGiverCareReceipentPairHandler(w http.ResponseWriter, r *http.Request) {

	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.GetMongoClient(mongoURI)
	if err != nil {
		w.Write([]byte("Unable to connect to mongo" + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	CareGiverCollection := client.Database("mydb").Collection("caregiver")
	CareReceipentCollection := client.Database("mydb").Collection("careReceipent")

	cr1 := models.CareReceipentModel{
		ID:                      uuid.NewString(),
		CapabilityPermissionMap: make(map[string]models.CapabilityPermissionSet),
	}

	cg1 := models.CareGiverModel{
		ID: uuid.NewString(),
	}

	cap1 := models.CapabilityPermissionSet{
		ID:          uuid.NewString(),
		Name:        "UserProfile",
		Permissions: []models.Permission{models.Create, models.Read},
	}

	cg1.CareReceipentIds = append(cg1.CareReceipentIds, cr1.ID)
	cr1.AuthorizedCareGiverIds = append(cr1.AuthorizedCareGiverIds, cg1.ID)
	cr1.CapabilityPermissionMap[cap1.ID] = cap1

	CareGiverCollection.InsertOne(context.Background(), cg1)
	CareReceipentCollection.InsertOne(context.Background(), cr1)

	w.Write(([]byte("Success")))
	w.WriteHeader(http.StatusOK)

}
