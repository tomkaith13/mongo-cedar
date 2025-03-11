package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tomkaith13/mongo-cedar/cedar_entity"
	"github.com/tomkaith13/mongo-cedar/models"
)

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody models.CheckRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eMap, err := cedar_entity.GenerateCareGiverEntity(reqBody.Principal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger := log.Default()
	b, err := eMap.MarshalJSON()
	logger.Printf("eMap: %s", string(b))

	// fetch care receipents doc to compose context
	w.WriteHeader(http.StatusOK)

}
