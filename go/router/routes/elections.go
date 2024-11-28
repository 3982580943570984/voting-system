package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"voting-system/services"

	"github.com/go-chi/chi/v5"
)

func ElectionsRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getAllElections)

	r.Get("/{id}", getElection)

	return r
}

func getAllElections(w http.ResponseWriter, r *http.Request) {
	elections, err := services.NewElections().
		GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(elections)
}

func getElection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid election ID", http.StatusBadRequest)
		return
	}

	election, err := services.NewElections().
		GetById(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(election)
}
