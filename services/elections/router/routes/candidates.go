package routes

import (
	"elections/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Candidates() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getCandidates)

	r.Post("/", createCandidate)

	r.Put("/", updateCandidate)

	r.Delete("/", deleteCandidate)

	return r
}

// @Summary Создать кандидата
// @Description Создает нового кандидата на основе данных, переданных в теле запроса.
// @Tags Кандидаты
// @Accept  json
// @Produce  json
// @Param id path int true "ID выборов" // Уникальный идентификатор выборов, к которым добавляется кандидат
// @Param candidate body services.CandidateCreate true "Данные для создания кандидата"
// @Success 201 {object} map[string]int "Успешно создан, возвращает идентификатор созданного кандидата"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id}/candidates [post]
func createCandidate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	cc := repositories.CandidateCreate{ElectionId: id}
	if err := json.NewDecoder(r.Body).Decode(&cc); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	candidate, err := repositories.NewCandidates().Create(r.Context(), &cc)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]int{"id": candidate.ID})
}

// @Summary Получить всех кандидатов
// @Description Эта функция возвращает список всех кандидатов для выбранных выборов по ID.
// @Tags Кандидаты
// @Accept json
// @Produce json
// @Param id path int true "ID выборов" // Уникальный идентификатор выборов
// @Success 200 {array} generated.Candidate "Список кандидатов"
// @Failure 400 {string} string "Неверный формат ID"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id}/candidates [get]
func getCandidates(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}

	candidates, err := repositories.NewCandidates().GetByElectionId(r.Context(), id)
	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(candidates)
}

// @Summary Обновить кандидата
// @Description Обновляет информацию о кандидате.
// @Tags Кандидаты
// @Accept json
// @Produce json
// @Param id path int true "ID выборов"
// @Param candidate body services.CandidateUpdate true "Данные для обновления кандидата"
// @Success 200 {object} map[string]int "Успешно обновлено, возвращает идентификатор обновленного кандидата"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 404 {string} string "Кандидат не найден"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id}/candidates [put]
func updateCandidate(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}

	var cu repositories.CandidateUpdate
	if err := json.NewDecoder(r.Body).Decode(&cu); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	candidate, err := repositories.NewCandidates().Update(r.Context(), &cu)
	if err != nil {
		http.Error(w, "Error during candidate update: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]int{"id": candidate.ID})
}

// @Summary Удалить кандидата
// @Description Удаляет кандидата.
// @Tags Кандидаты
// @Accept json
// @Produce json
// @Param id path int true "ID выборов"
// @Param candidate body services.CandidateDelete true "Данные для удаления кандидата"
// @Success 204 {string} string "Кандидат успешно удален"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 404 {string} string "Кандидат не найден"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id}/candidates [delete]
func deleteCandidate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID"+err.Error(), http.StatusBadRequest)
		return
	}

	cd := repositories.CandidateDelete{ElectionId: id}
	if err := json.NewDecoder(r.Body).Decode(&cd); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = repositories.NewCandidates().Delete(r.Context(), &cd)
	if err != nil {
		http.Error(w, "Error during candidate deletion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
