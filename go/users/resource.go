package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Resource struct {
	Storage *Storage
}

func NewResource(s *Storage) *Resource {
	return &Resource{Storage: s}
}

func (rs Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.GetAll)
	r.Post("/", rs.Post)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Get)
		r.Put("/", rs.Put)
		r.Delete("/", rs.Delete)
	})

	return r
}

// curl localhost:3000/users --request POST --data '{"email":"example@example.com", "password":"password", "is_active":true}'
func (rs Resource) Post(w http.ResponseWriter, r *http.Request) {
	var user User

	error := json.NewDecoder(r.Body).Decode(&user)

	if error != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID, error := rs.Storage.Create(user)

	if error != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]int{"user_id": userID})
}

// curl localhost:3000/users/1 --request GET
func (rs Resource) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, error := strconv.Atoi(idStr)

	if error != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, error := rs.Storage.GetByID(id)

	if error != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

// curl localhost:3000/users --request GET
func (rs Resource) GetAll(w http.ResponseWriter, r *http.Request) {
	usersList, error := rs.Storage.GetAll()

	if error != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(usersList)
}

// curl localhost:3000/users/1 --request PUT --data '{"email":"example@example.com", "password":"password", "is_active":true}'
func (rs Resource) Put(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, error := strconv.Atoi(idStr)

	if error != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user User

	error = json.NewDecoder(r.Body).Decode(&user)

	if error != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user.Id = id

	error = rs.Storage.Update(user)

	if error != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// curl localhost:3000/users/1 --request DELETE
func (rs Resource) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, error := strconv.Atoi(idStr)

	if error != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	error = rs.Storage.Delete(id)

	if error != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
