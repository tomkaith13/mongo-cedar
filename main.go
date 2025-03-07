package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tomkaith13/mongo-cedar/handler"
)

func main() {
	r := chi.NewRouter()

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Post("/insert-example", handler.CreateCareGiverCareReceipentPairHandler)

	http.ListenAndServe(":8888", r)
}
