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

	http.ListenAndServe(":8888", r)
}
