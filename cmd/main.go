package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"mymodule/pkg/db"
	"mymodule/pkg/handlers"
)

func main() {
	r := chi.NewRouter()
	DB := db.Init()
	h := handlers.New(DB)
	r.Use(middleware.Logger)

	r.HandleFunc("/addBook", h.Store)
	r.HandleFunc("/getBook", h.GetAllBooks)
	r.HandleFunc("/addReview", h.AddReview)
	r.HandleFunc("/getReview", h.GetReview)

	http.ListenAndServe(":3000", r)
}
