package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"mymodule/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) GetReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	log.Println(id)
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var data []models.Book
	json.Unmarshal(body, &data)

	// var reviews []models.Review
	// for _, books := range data {
	// 	if books.ID == id {
	// 		reviews = append(reviews, books.Review...)
	// 	}
	// }

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("reviews")
}
