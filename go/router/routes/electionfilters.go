package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"voting-system/services"

	"github.com/go-chi/chi/v5"
)

func ElectionFiltersRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getFilters)

	r.Put("/", updateFilters)

	return r
}

func getFilters(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}

	settings, err := services.NewElections().GetSettings(r.Context(), id)
	if err != nil {
		http.Error(w, "Error during filters retrieval"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(settings)
}

// TODO: implement
func updateFilters(w http.ResponseWriter, r *http.Request) {

}
