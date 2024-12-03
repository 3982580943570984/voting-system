package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"voting-system/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func ProfileRoutes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token), jwtauth.Authenticator(token))

		r.Get("/", getProfile)

		r.Put("/", updateProfile)
	})

	return r
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	id, err := RetrieveIdFromToken(r.Context())

	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	profile, err := services.NewProfiles().GetByUserId(r.Context(), id)

	if err != nil {
		http.Error(w, "Error retrieving profile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(profile)
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	id, err := RetrieveIdFromToken(r.Context())

	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	pu := services.ProfileUpdate{
		UserID: id,
	}

	if err := json.NewDecoder(r.Body).Decode(&pu); err != nil {
		log.Print(err)

		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = services.NewProfiles().Update(r.Context(), &pu)

	if err != nil {
		http.Error(w, "Error updating profile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
