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

	r.Get("/", getAll)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token), jwtauth.Authenticator(token))

		r.Get("/filtered", getFiltered)

		r.Get("/created", getCreated)

		r.Post("/", create)
	})

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", get)

		r.Put("/", update)

		r.Delete("/", delete)

		r.Mount("/candidates", CandidatesRoutes())

		r.Mount("/tags", TagsRoutes())

		r.Mount("/settings", ElectionSettingsRoutes())

		r.Mount("/comments", CommentsRoutes())
	})

	return r
}

// @Summary Создать выборы
// @Description Создает новые выборы.
// @Tags Выборы
// @Accept  json
// @Produce  json
// @Param election body services.ElectionCreate true "Данные для создания выборов"
// @Success 201 {object} map[string]int "Успешно создано, возвращает идентификатор созданных выборов"
// @Failure 400 {string} string "Неверный формат запроса или данные в теле запроса"
// @Failure 401 {string} string "Неавторизованный доступ"
// @Failure 403 {string} string "Доступ запрещен"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections [post]
// @Security Bearer
func create(w http.ResponseWriter, r *http.Request) {
	userID, err := RetrieveIdFromToken(r.Context())
	if err != nil {
		http.Error(w, "Invalid authentication token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	payload := services.ElectionCreate{UserID: userID}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.NewUsers().GetById(r.Context(), payload.UserID)
	if err != nil {
		http.Error(w, "Error fetching user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !user.IsOrganizer {
		http.Error(w, "Only organizers can create elections", http.StatusForbidden)
		return
	}

	election, err := services.NewElections().Create(r.Context(), &payload)
	if err != nil {
		http.Error(w, "Error creating election with candidates: "+err.Error(), http.StatusInternalServerError)
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
func getAll(w http.ResponseWriter, r *http.Request) {
	elections, err := services.NewElections().
		GetAllWithDuration(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(elections)
}

// @Summary Получить отфильтрованные выборы
// @Description Возвращает список выборов, отфильтрованных по определенным критериям, доступных текущему пользователю.
// @Tags Выборы
// @Accept json
// @Produce json
// @Success 200 {array} generated.Election "Отфильтрованный список выборов"
// @Failure 401 {object} map[string]string "Неавторизованный доступ"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /elections/filtered [get]
// @Security Bearer
func getFiltered(w http.ResponseWriter, r *http.Request) {
	userID, err := RetrieveIdFromToken(r.Context())
	if err != nil {
		http.Error(w, "Invalid authentication token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	elections, err := services.NewElections().GetAllFiltered(r.Context(), userID)
	if err != nil {
		http.Error(w, "Error fetching filtered elections: "+err.Error(), http.StatusInternalServerError)
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
func get(w http.ResponseWriter, r *http.Request) {
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

// @Summary Получить выборы, созданные пользователем
// @Description Эта функция возвращает список выборов, созданных текущим пользователем.
// @Tags Выборы
// @Accept json
// @Produce json
// @Success 200 {array} generated.Election "Список выборов, созданных пользователем"
// @Failure 401 {object} map[string]string "Неавторизованный доступ"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /elections/created [get]
// @Security Bearer
func getCreated(w http.ResponseWriter, r *http.Request) {
	userID, err := RetrieveIdFromToken(r.Context())
	if err != nil {
		http.Error(w, "Invalid authentication token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	elections, err := services.NewElections().GetByUserId(r.Context(), userID)

	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(elections)
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
func update(w http.ResponseWriter, r *http.Request) {
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
func delete(w http.ResponseWriter, r *http.Request) {
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
