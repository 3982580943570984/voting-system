package routes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"voting-system/ent/generated"
	"voting-system/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func ElectionsRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getAllElections)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token), jwtauth.Authenticator(token))

		r.Post("/", createElection)
	})

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", getElection)

		r.Put("/", updateElection)

		r.Delete("/", deleteElection)

		r.Mount("/candidates", CandidatesRoutes())

		r.Mount("/tags", TagsRoutes())

		r.Mount("/settings", ElectionSettingsRoutes())

		r.Mount("/comments", CommentsRoutes())
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
// @Security Bearer
func createElection(w http.ResponseWriter, r *http.Request) {
	id, err := RetrieveIdFromToken(r.Context())

	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	ec := services.ElectionCreate{UserID: id}

	if err := json.NewDecoder(r.Body).Decode(&ec); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.NewUsers().GetById(r.Context(), ec.UserID)

	if err != nil {
		http.Error(w, "Error finding user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !user.IsOrganizer {
		http.Error(w, "Not organizer can't create election", http.StatusBadRequest)
		return
	}

	election, err := services.NewElections().Create(r.Context(), &ec)

	if err != nil {
		http.Error(w, "Error creating election: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]int{"id": election.ID})
}

// @Summary Получить все выборы
// @Description Эта функция возвращает список всех доступных выборов.
// @Tags Выборы
// @Accept json
// @Produce json
// @Success 200 {array} generated.Election "Список выборов"
// @Failure 400 {string} string "Ошибка запроса"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections [get]
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

// @Summary Получить выборы по ID
// @Description Эта функция возвращает информацию о конкретных выборах по ID.
// @Tags Выборы
// @Accept json
// @Produce json
// @Param id path int true "ID выборов"
// @Success 200 {object} generated.Election "Информация о выборах"
// @Failure 400 {string} string "Неверный формат ID"
// @Failure 404 {string} string "Выборы не найдены"
// @Router /elections/{id} [get]
func getElection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid election ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	election, err := services.NewElections().GetById(r.Context(), id)

	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(election)
}

// @Summary Обновить выборы
// @Description Эта функция обновляет данные выборов по ID.
// @Tags Выборы
// @Accept json
// @Produce json
// @Param id path int true "ID выборов для обновления"
// @Param election body services.ElectionUpdate true "Данные для обновления выборов"
// @Success 200 {object} map[string]int "Успешно обновлено, возвращает идентификатор обновленных выборов"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 404 {string} string "Выборы с указанным ID не найдены"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id} [put]
func updateElection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}

	eu := services.ElectionUpdate{
		ID: id,
	}

	if err := json.NewDecoder(r.Body).Decode(&eu); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusInternalServerError)
		return
	}

	election, err := services.NewElections().Update(r.Context(), &eu)

	if err != nil {
		http.Error(w, "Error during election update: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]int{"id": election.ID})
}

// @Summary Удалить выборы
// @Description Эта функция удаляет выборы по ID.
// @Tags Выборы
// @Accept json
// @Produce json
// @Param id path int true "ID выборов для удаления"
// @Success 204 {string} string "Успешно удалено, нет контента"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 404 {string} string "Выборы с указанным ID не найдены"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id} [delete]
func deleteElection(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}

	err = services.NewElections().DeleteById(r.Context(), id)

	if err != nil {
		if errors.Is(err, &generated.NotFoundError{}) {
			http.Error(w, fmt.Sprintf("Election with ID %d not found", id), http.StatusNotFound)
			return
		}

		http.Error(w, "Error during election deletion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func RetrieveIdFromToken(ctx context.Context) (int, error) {
	_, claims, err := jwtauth.FromContext(ctx)

	if err != nil {
		return -1, err
	}

	idFloat, ok := claims["id"].(float64)

	if !ok {
		return -1, errors.New("invalid or missing 'id' in token claims")
	}

	return int(idFloat), nil
}
