package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/tomkaith13/mongo-cedar/models"
	"github.com/tomkaith13/mongo-cedar/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func CreateCareGiverCareReceipentPairHandler(w http.ResponseWriter, r *http.Request) {

	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.GetMongoClient(mongoURI)
	if err != nil {
		w.Write([]byte("Unable to connect to mongo" + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	CareGiverCollection := client.Database("mydb").Collection("caregivers")
	CareReceipentCollection := client.Database("mydb").Collection("carereceipents")

	// transaction options for cluster
	wc := writeconcern.Majority()
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	session, err := client.StartSession()
	if err != nil {
		w.Write([]byte("Unable to start session"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer session.EndSession(context.Background())

	err = session.StartTransaction(txnOpts)
	if err != nil {
		defer session.AbortTransaction(context.Background())
		w.Write([]byte("Unable to start transaction"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cr1 := models.CareReceipentModel{
		ID:                               "cr1",
		FirstName:                        gofakeit.FirstName(),
		LastName:                         gofakeit.LastName(),
		Email:                            gofakeit.Email(),
		AuthorizedCareGiverIds:           make(map[string]bool),
		CareGiverCapabilityPermissionMap: make(map[string]map[string]models.CapabilityPermissionSet),
	}

	cg1 := models.CareGiverModel{
		ID:                     "cg1",
		FirstName:              gofakeit.FirstName(),
		LastName:               gofakeit.LastName(),
		Email:                  gofakeit.Email(),
		Phone:                  gofakeit.Phone(),
		CareReceipentIds:       make(map[string]bool),
		CareReceipentInviteMap: make(map[string]models.InviteStatus),
		AllowedResourceIds:     make(map[string]bool),
	}

	cap1 := models.CapabilityPermissionSet{
		ID:          "UserProfile",
		Name:        "UserProfile",
		Permissions: []models.Permission{models.Create, models.Read},
	}
	cap2 := models.CapabilityPermissionSet{
		ID:          "Documents",
		Name:        "Documents",
		Permissions: []models.Permission{models.Create, models.Read, models.Delete},
	}

	cg1.CareReceipentIds[cr1.ID] = true
	cg1.CareReceipentInviteMap[cr1.ID] = models.Accepted
	cg1.AllowedResourceIds["HomePage"] = true

	cr1.AuthorizedCareGiverIds[cg1.ID] = true

	cr1.CareGiverCapabilityPermissionMap[cg1.ID] = make(map[string]models.CapabilityPermissionSet)
	cr1.CareGiverCapabilityPermissionMap[cg1.ID][cap1.ID] = cap1
	cr1.CareGiverCapabilityPermissionMap[cg1.ID][cap2.ID] = cap2

	CareGiverCollection.InsertOne(context.Background(), cg1)
	CareReceipentCollection.InsertOne(context.Background(), cr1)

	err = session.CommitTransaction(context.Background())
	if err != nil {
		defer session.AbortTransaction(context.Background())
		w.Write([]byte("Unable to commit transaction"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(([]byte(fmt.Sprintf("caregiver: %s", cg1.ID))))
	w.WriteHeader(http.StatusOK)

}
