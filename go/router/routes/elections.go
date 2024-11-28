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

	r.Route("/", func(r chi.Router) {
		r.Post("/", createElection)

		r.Get("/", getAllElections)

		r.Get("/{id}", getElection)
	})

	return r
}

// @Summary Создать выборы
// @Description Создает новые выборы на основе данных, переданных в теле запроса.
// @Tags Выборы
// @Accept  json
// @Produce  json
// @Param election body services.ElectionCreate true "Данные для создания выборов"
// @Success 201 {object} int "Успешно создано, возвращает идентификатор созданных выборов"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections [post]
func createElection(w http.ResponseWriter, r *http.Request) {
	var ec services.ElectionCreate

	if err := json.NewDecoder(r.Body).Decode(&ec); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.NewUsers().GetById(r.Context(), ec.UserID)

	if err != nil {
		http.Error(w, "Error finding user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	role, err := user.QueryRole().Only(r.Context())

	if err != nil {
		http.Error(w, "Error querying user's role: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if role.IsVoter {
		http.Error(w, "Voter can't create election", http.StatusBadRequest)
		return
	}

	election, err := services.NewElections().
		Create(r.Context(), &ec)

	if err != nil {
		http.Error(w, "Error creating election: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]int{"id": election.ID})
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
