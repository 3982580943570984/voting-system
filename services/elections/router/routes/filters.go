package routes

import (
	"elections/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Filters() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getFilters)

	r.Put("/", updateFilters)

	return r
}

// @Summary Получить фильтры выборов
// @Description Возвращает фильтры для указанных выборов.
// @Tags Фильтры
// @Accept json
// @Produce json
// @Param id path int true "ID выборов"
// @Success 200 {object} generated.ElectionFilters "Фильтры выборов"
// @Failure 400 {string} string "Неверный ID выборов"
// @Failure 404 {string} string "Выборы не найдены"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id}/filters [get]
func getFilters(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}

	filters, err := repositories.NewElections().GetFilters(r.Context(), id)
	if err != nil {
		http.Error(w, "Error during filters retrieval"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(filters)
}

// @Summary Обновить фильтры выборов
// @Description Обновляет фильтры для указанных выборов.
// @Tags Фильтры
// @Accept json
// @Produce json
// @Param id path int true "ID выборов"
// @Param filters body services.ElectionFiltersUpdate true "Обновленные данные фильтров"
// @Success 204 {string} string "Фильтры успешно обновлены"
// @Failure 400 {string} string "Неверный ID выборов или входные данные"
// @Failure 404 {string} string "Выборы не найдены"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id}/filters [put]
func updateFilters(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}

	efu := repositories.FiltersUpdate{ElectionID: id}
	if err := json.NewDecoder(r.Body).Decode(&efu); err != nil {
		http.Error(w, "Invalid input"+err.Error(), http.StatusBadRequest)
		return
	}

	err = repositories.NewFilters().Update(r.Context(), &efu)
	if err != nil {
		http.Error(w, "Error during filters update: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
