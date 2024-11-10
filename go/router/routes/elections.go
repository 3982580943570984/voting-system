package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"voting-system/database"
	"voting-system/ent"

	"github.com/go-chi/chi/v5"
)

func ElectionsRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getAllElections)
	r.Post("/", createElection)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", getElection)
		r.Put("/", updateElection)
		r.Delete("/", deleteElection)
	})

	return r
}

// Пример запроса для создания Election:
// curl localhost:3000/elections --request POST --data '{"title":"Выборы президента", "description":"Описание выборов", "start_date":"2024-05-01T00:00:00Z", "end_date":"2024-05-10T00:00:00Z", "is_active":true}'
func createElection(w http.ResponseWriter, r *http.Request) {
	var election ent.Election

	if err := json.NewDecoder(r.Body).Decode(&election); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	newElection, err := database.Elections.
		Create().
		SetTitle(election.Title).
		SetDescription(election.Description).
		SetStartDate(election.StartDate).
		SetEndDate(election.EndDate).
		SetIsActive(election.IsActive).
		Save(context.Background())

	if err != nil {
		http.Error(w, "Не удалось создать выборы", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": newElection.ID})
}

// curl localhost:3000/elections/1 --request GET
func getElection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Неверный ID выборов", http.StatusBadRequest)
		return
	}

	election, err := database.Elections.Get(context.Background(), id)
	if err != nil {
		http.Error(w, "Не удалось получить выборы", http.StatusInternalServerError)
		return
	}

	if election == nil {
		http.Error(w, "Выборы не найдены", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(election)
}

// curl localhost:3000/elections --request GET
func getAllElections(w http.ResponseWriter, r *http.Request) {
	elections, err := database.Elections.Query().All(context.Background())
	if err != nil {
		http.Error(w, "Не удалось получить список выборов", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(elections)
}

// curl localhost:3000/elections/1 --request PUT --data '{"title":"Обновленные выборы", "description":"Новое описание", "start_date":"2024-06-01T00:00:00Z", "end_date":"2024-06-10T00:00:00Z", "is_active":false}'
func updateElection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Неверный ID выборов", http.StatusBadRequest)
		return
	}

	var election ent.Election
	if err := json.NewDecoder(r.Body).Decode(&election); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	updatedElection, err := database.Elections.
		UpdateOneID(id).
		SetTitle(election.Title).
		SetDescription(election.Description).
		SetStartDate(election.StartDate).
		SetEndDate(election.EndDate).
		SetIsActive(election.IsActive).
		Save(context.Background())

	if err != nil {
		http.Error(w, "Не удалось обновить выборы", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": updatedElection.ID})
}

// Пример запроса для удаления Election:
// curl localhost:3000/elections/1 --request DELETE
func deleteElection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Неверный ID выборов", http.StatusBadRequest)
		return
	}

	if err := database.Elections.DeleteOneID(id).Exec(context.Background()); err != nil {
		http.Error(w, "Не удалось удалить выборы", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}
