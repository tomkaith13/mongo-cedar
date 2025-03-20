package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cedar-policy/cedar-go"
	"github.com/tomkaith13/mongo-cedar/cedar_context"
	"github.com/tomkaith13/mongo-cedar/cedar_entity"
	"github.com/tomkaith13/mongo-cedar/cedar_policy"
	"github.com/tomkaith13/mongo-cedar/models"
	"github.com/tomkaith13/mongo-cedar/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody models.CheckRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.GetMongoClient(mongoURI)
	if err != nil {
		w.Write([]byte("Unable to connect to mongo" + err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

	eMap, err := cedar_entity.GenerateCareGiverEntity(reqBody.CareGiverId, reqBody.CareReceipentId)
	if err != nil {
		w.Write(fmt.Append(nil, "Authorized: false"))
		return
	}
	eMap, err = cedar_entity.AddActionEntity(reqBody.Action, eMap)
	if err != nil {
		w.Write(fmt.Append(nil, "Authorized: false"))
		return
	}

	eMap, err = cedar_entity.AddResourceEntity(reqBody.Resource, eMap)
	if err != nil {
		w.Write(fmt.Append(nil, "Authorized: false"))
		return
	}

	b, _ := eMap.MarshalJSON()
	fmt.Printf("eMap: %s", string(b))

	// fetch care receipents doc to compose context
	cedarCtx, err := cedar_context.GenerateContext(reqBody.CareReceipentId, reqBody.CareGiverId, reqBody.Resource)
	if err != nil {
		fmt.Printf("Error generating context: %s.\nrejecting authz request", err.Error())
		w.Write(fmt.Append(nil, "Authorized: false"))
		return
	}
	b, _ = cedarCtx.MarshalJSON()
	fmt.Printf("Context: %s", string(b))
	err = session.CommitTransaction(context.Background())
	if err != nil {
		defer session.AbortTransaction(context.Background())
		w.Write([]byte("Unable to commit transaction"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	request := cedar.Request{
		Principal: cedar.NewEntityUID("CareGiver", cedar.String(reqBody.CareGiverId)),
		Action:    cedar.NewEntityUID("Action", cedar.String(reqBody.Action)),
		Resource:  cedar.NewEntityUID("Capability", cedar.String(reqBody.Resource)),
		Context:   *cedarCtx,
	}

	ok, diag := cedar_policy.PolicySet.IsAuthorized(eMap, request)
	fmt.Printf("Is Authorized: %t", ok)
	fmt.Printf("Diagnostic: %+v", diag)

	w.Write(fmt.Appendf(nil, "Authorized: %t", ok))

	w.WriteHeader(http.StatusOK)

}

func testPolicies() {

	b, err := os.ReadFile("./../policy.cedar")
	if err != nil {
		return
	}

	PolicySet, err := cedar.NewPolicySetFromBytes("./../policy.cedar", b)
	if err != nil {
		return
	}

	b, err = PolicySet.MarshalJSON()
	if err != nil {
		return
	}
	fmt.Printf("policies: %s\n", string(b))

	const entitiesJSON = `[
  {
    "uid": { "type": "CareGiver", "id": "cg1" },
    "attrs": { "OwnResourceSet": ["HomePage"] },
    "parents": []
  },
  {
    "uid": { "type": "Capability", "id": "HomePage" },
    "attrs": { "name": "HomePage" },
    "parents": []
  }
]`

	var entities cedar.EntityMap
	if err := json.Unmarshal([]byte(entitiesJSON), &entities); err != nil {
		log.Fatal(err)
	}

	req := cedar.Request{
		Principal: cedar.NewEntityUID("CareGiver", "cg1"),
		Action:    cedar.NewEntityUID("Action", "view"),
		Resource:  cedar.NewEntityUID("Capability", "HomePage"),
		Context: cedar.NewRecord(cedar.RecordMap{
			"impersonation": cedar.False,
		}),
	}

	fmt.Println("------start--------")

	ok, diag := PolicySet.IsAuthorized(entities, req)
	fmt.Println(ok)
	fmt.Println(diag)

}
