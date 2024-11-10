package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"voting-system/database"

	"github.com/go-chi/chi/v5"
)

type userInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
}

func UsersRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getAllUsers)
	r.Post("/", createUser)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", getUser)
		r.Put("/", updateUser)
		r.Delete("/", deleteUser)
	})

	return r
}

// curl localhost:3000/users --request POST --data '{"email":"example@example.com", "password":"password", "is_active":true}'
func createUser(w http.ResponseWriter, r *http.Request) {
	var input userInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user := database.Users.
		Create().
		SetEmail(input.Email).
		SetPassword(input.Password).
		SetIsActive(input.IsActive).
		SaveX(r.Context())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]int{"id": user.ID})
}

// curl localhost:3000/users/1 --request GET
func getUser(w http.ResponseWriter, r *http.Request) {
	id, error := strconv.Atoi(chi.URLParam(r, "id"))

	if error != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user := database.Users.GetX(r.Context(), id)

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

// curl localhost:3000/users --request GET
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := database.Users.
		Query().
		AllX(r.Context())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

// curl localhost:3000/users/1 --request PUT --data '{"email":"example@example.com", "password":"password", "is_active":true}'
func updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var input userInput

	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedUser := database.Users.
		UpdateOneID(id).
		SetEmail(input.Email).
		SetPassword(input.Password).
		SetIsActive(input.IsActive).
		SaveX(r.Context())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]int{"id": updatedUser.ID})
}

// curl localhost:3000/users/1 --request DELETE
func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	database.Users.
		DeleteOneID(id).
		ExecX(r.Context())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]int{"id": id})
}
