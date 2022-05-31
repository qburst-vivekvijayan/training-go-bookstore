package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/qburst-vivekvijayan/training-go-bookstore.git/pkg/db"
	"github.com/qburst-vivekvijayan/training-go-bookstore.git/pkg/handlers"
)

func main() {
	r := chi.NewRouter()
	DB := db.Init()
	h := handlers.New(DB)
	r.Use(middleware.Logger)

	r.HandleFunc("/store", h.Store)

	http.ListenAndServe(":3000", r)
}
