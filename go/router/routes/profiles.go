package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"voting-system/database"
	"voting-system/ent"

	"github.com/go-chi/chi/v5"
)

type profileInput struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthdate time.Time `json:"birthdate,omitempty"`
	UserID    int       `json:"user_id"`
}

func ProfilesRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", getAllProfiles)
	router.Post("/", createProfile)

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", getProfile)
		r.Put("/", updateProfile)
		r.Delete("/", deleteProfile)
	})

	return router
}

func createProfile(w http.ResponseWriter, r *http.Request) {
	var input profileInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user := database.Users.GetX(context.Background(), input.UserID)

	profile := database.Profiles.
		Create().
		SetFirstName(input.FirstName).
		SetLastName(input.LastName).
		SetBirthdate(input.Birthdate).
		SetUser(user).
		SaveX(r.Context())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]int{"id": profile.ID})
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	profile := database.Profiles.GetX(r.Context(), id)

	if profile == nil {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(profile)
}

func getAllProfiles(w http.ResponseWriter, r *http.Request) {
	profiles := database.Profiles.
		Query().
		WithUser().
		AllX(context.Background())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(profiles)
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	var input struct {
		FirstName string    `json:"first_name,omitempty"`
		LastName  string    `json:"last_name,omitempty"`
		Birthdate time.Time `json:"birthdate,omitempty"`
		UserID    int       `json:"user_id,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	update := database.Profiles.UpdateOneID(id)

	if input.FirstName != "" {
		update.SetFirstName(input.FirstName)
	}
	if input.LastName != "" {
		update.SetLastName(input.LastName)
	}
	// Only set Birthdate if it's provided (zero value is ignored)
	if !input.Birthdate.IsZero() {
		update.SetBirthdate(input.Birthdate)
	}
	if input.UserID != 0 {
		// Ensure the associated User exists
		user, err := database.Users.Get(context.Background(), input.UserID)
		if err != nil {
			http.Error(w, "User not found", http.StatusBadRequest)
			return
		}
		update.SetUser(user)
	}

	updatedProfile, err := update.Save(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			http.Error(w, "Profile not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": updatedProfile.ID})
}

// deleteProfile deletes a Profile by its ID.
func deleteProfile(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	err = database.Profiles.
		DeleteOneID(id).
		Exec(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			http.Error(w, "Profile not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}
