package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"voting-system/services"

	"github.com/go-chi/chi/v5"
)

func ElectionSettingsRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getSettings)

	r.Put("/", updateSettings)

	return r
}

// @Summary Получить настройки выборов
// @Description Возвращает настройки выборов по ID.
// @Tags Настройки выборов
// @Accept json
// @Produce json
// @Param id path int true "ID выборов"
// @Success 200 {object} generated.ElectionSettings "Настройки выборов"
// @Failure 400 {string} string "Неверный формат ID"
// @Failure 404 {string} string "Выборы не найдены"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id}/settings [get]
func getSettings(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}

	settings, err := services.NewElections().GetSettings(r.Context(), id)

	if err != nil {
		http.Error(w, "Error during settings retrieval"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(settings)
}

// TODO: implement logic
func updateSettings(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}
}
