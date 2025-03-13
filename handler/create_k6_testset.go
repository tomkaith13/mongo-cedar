package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/tomkaith13/mongo-cedar/models"
	"github.com/tomkaith13/mongo-cedar/mongo"
)

const NumOfCGs = 100
const NumOfCRs = 10000

func CreatePerfTestSetHandler(w http.ResponseWriter, r *http.Request) {

	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.GetMongoClient(mongoURI)
	if err != nil {
		w.Write([]byte("Unable to connect to mongo" + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	CareGiverCollection := client.Database("mydb").Collection("caregivers")
	CareReceipentCollection := client.Database("mydb").Collection("carereceipents")

	for i := 1; i <= NumOfCGs; i++ {

		cg := models.CareGiverModel{
			ID:               "cg" + strconv.Itoa(i),
			FirstName:        gofakeit.FirstName(),
			LastName:         gofakeit.LastName(),
			Email:            gofakeit.Email(),
			Phone:            gofakeit.Phone(),
			CareReceipentIds: make(map[string]bool),
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

		for j := 1; j <= NumOfCRs; j++ {
			cr := models.CareReceipentModel{
				ID:                               "cr" + strconv.Itoa(i) + "::" + strconv.Itoa(j),
				FirstName:                        gofakeit.FirstName(),
				LastName:                         gofakeit.LastName(),
				Email:                            gofakeit.Email(),
				AuthorizedCareGiverIds:           make(map[string]bool),
				CareGiverCapabilityPermissionMap: make(map[string]map[string]models.CapabilityPermissionSet),
			}
			cg.CareReceipentIds[cr.ID] = true
			cr.AuthorizedCareGiverIds[cg.ID] = true

			cr.CareGiverCapabilityPermissionMap[cg.ID] = make(map[string]models.CapabilityPermissionSet)
			cr.CareGiverCapabilityPermissionMap[cg.ID][cap1.ID] = cap1
			cr.CareGiverCapabilityPermissionMap[cg.ID][cap2.ID] = cap2

			_, err := CareReceipentCollection.InsertOne(context.Background(), cr)
			if err != nil {
				fmt.Printf("cr insert of id: %s did not work. error: %s\n", cr.ID, err)
			}
		}

		_, err := CareGiverCollection.InsertOne(context.Background(), cg)
		if err != nil {
			fmt.Printf("cg insert of id: %s did not work. error got: %s\n", cg.ID, err)
		}

	}

}
