package routes

import (
	"elections/repositories"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Settings() chi.Router {
	r := chi.NewRouter()

	r.Get("/", get)

	r.Put("/", update)

	return r
}

// @Summary Получить настройки выборов
// @Description Возвращает настройки выборов по ID.
// @Tags Настройки
// @Accept json
// @Produce json
// @Param id path int true "ID выборов"
// @Success 200 {object} generated.ElectionSettings "Настройки выборов"
// @Failure 400 {string} string "Неверный формат ID"
// @Failure 404 {string} string "Выборы не найдены"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id}/settings [get]
func get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}

	settings, err := repositories.NewElections().GetSettings(r.Context(), id)
	if err != nil {
		http.Error(w, "Error during settings retrieval"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(settings)
}

// @Summary Обновить настройки выборов
// @Description Обновляет настройки для указанных выборов.
// @Tags Настройки
// @Accept json
// @Produce json
// @Param id path int true "ID выборов"
// @Param settings body services.ElectionSettingsUpdate true "Обновленные данные настроек"
// @Success 204 {string} string "Настройки успешно обновлены"
// @Failure 400 {string} string "Неверный ID выборов или входные данные"
// @Failure 404 {string} string "Выборы не найдены"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id}/settings [put]
func update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}

	esu := repositories.ElectionSettingsUpdate{ElectionID: id}
	if err := json.NewDecoder(r.Body).Decode(&esu); err != nil {
		log.Print(err.Error())
		http.Error(w, "Invalid input"+err.Error(), http.StatusBadRequest)
		return
	}

	err = repositories.NewElectionSettings().Update(r.Context(), &esu)
	if err != nil {
		http.Error(w, "Error during settings update: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
