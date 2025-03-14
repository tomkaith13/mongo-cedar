package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cedar-policy/cedar-go"
	"github.com/tomkaith13/mongo-cedar/cedar_entity"
	"github.com/tomkaith13/mongo-cedar/cedar_policy"
	"github.com/tomkaith13/mongo-cedar/models"
)

func CheckSelfHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody models.CheckRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	eMap, err := cedar_entity.GenerateCareGiverEntity(reqBody.CareGiverId, reqBody.CareReceipentId)
	if err != nil {
		w.Write(fmt.Append(nil, "Authorized: false"))
		return
	}
	// eMap, err = cedar_entity.AddActionEntity(reqBody.Action, eMap)
	// if err != nil {
	// 	w.Write(fmt.Append(nil, "Authorized: false"))
	// 	return
	// }

	eMap, err = cedar_entity.AddResourceEntity(reqBody.Resource, eMap)
	if err != nil {
		w.Write(fmt.Append(nil, "Authorized: false"))
		return
	}

	logger := log.Default()
	b, _ := eMap.MarshalJSON()
	logger.Printf("eMap: %s", string(b))

	rm := cedar.RecordMap{}
	rm["impersonation"] = cedar.False
	record := cedar.NewRecord(rm)

	b, _ = record.MarshalJSON()

	logger.Printf("context: %s\n", string(b))

	request := cedar.Request{
		Principal: cedar.NewEntityUID("CareGiver", cedar.String(reqBody.CareGiverId)),
		Action:    cedar.NewEntityUID("Action", cedar.String(reqBody.Action)),
		Resource:  cedar.NewEntityUID("Capability", cedar.String(reqBody.Resource)),
		Context:   record,
	}

	ok, diag := cedar_policy.PolicySet.IsAuthorized(eMap, request)
	logger.Printf("Is Authorized: %t", ok)
	logger.Printf("Diagnostic: %s", diag)

	w.Write(fmt.Appendf(nil, "Authorized: %t", ok))

	w.WriteHeader(http.StatusOK)

}
