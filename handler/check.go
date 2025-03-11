package handler

import (
	"encoding/json"
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

	_, _ = cedar_entity.GenerateCareGiverEntity(reqBody.Principal)
	w.WriteHeader(http.StatusOK)

}
