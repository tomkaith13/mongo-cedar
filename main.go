package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tomkaith13/mongo-cedar/cedar_policy"
	"github.com/tomkaith13/mongo-cedar/handler"
)

func main() {
	r := chi.NewRouter()

	err := cedar_policy.LoadPolicy()
	if err != nil {
		panic(err)
	}

	p, err := cedar_policy.Policy.MarshalJSON()

	if err != nil {
		panic(err)
	}

	logger := log.Default()
	logger.Printf("Policy: %s\n", p)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Post("/insert-example", handler.CreateCareGiverCareReceipentPairHandler)
	r.Post("/insert-perf-data", handler.CreatePerfTestSetHandler)
	r.Post("/check", handler.CheckHandler)
	r.Post("/check-self", handler.CheckSelfHandler)

	// testingCedar()

	http.ListenAndServe(":8888", r)
}

// func testingCedar() {
// 	const policyCedar = `
// permit (principal == User::"bob",action,resource);
// permit (
// 	principal == User::"alice",
// 	action == Action::"view",
// 	resource in Album::"jane_vacation"
//   );

// `

// 	const entitiesJSON = `[
//   {
//     "uid": { "type": "User", "id": "alice" },
//     "attrs": { "age": 18 },
//     "parents": []
//   },
//   {
//     "uid": { "type": "User", "id": "bob" },
//     "attrs": { "age": 18 },
//     "parents": []
//   },
//   {
//     "uid": { "type": "Photo", "id": "VacationPhoto94.jpg" },
//     "attrs": {},
//     "parents": [{ "type": "Album", "id": "jane_vacation" }]
//   }
// ]`

// 	var policy cedar.Policy
// 	if err := policy.UnmarshalCedar([]byte(policyCedar)); err != nil {
// 		log.Fatal(err)
// 	}

// 	ps := cedar.NewPolicySet()
// 	ps.Add("policy0", &policy)

// 	var entities cedar.EntityMap
// 	if err := json.Unmarshal([]byte(entitiesJSON), &entities); err != nil {
// 		log.Fatal(err)
// 	}

// 	req := cedar.Request{
// 		Principal: cedar.NewEntityUID("User", "bob"),
// 		Action:    cedar.NewEntityUID("Action", "view"),
// 		Resource:  cedar.NewEntityUID("Photo", "VacationPhoto94.jpg"),
// 		Context: cedar.NewRecord(cedar.RecordMap{
// 			"demoRequest": cedar.True,
// 		}),
// 	}

// 	b, _ := req.Context.MarshalJSON()
// 	fmt.Println(string(b))

// 	ok, diag := ps.IsAuthorized(entities, req)
// 	fmt.Println(ok)
// 	fmt.Println(diag)

// }
