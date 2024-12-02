package routes

import (
	"encoding/json"
	"net/http"
	"voting-system/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func VotesRoutes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token), jwtauth.Authenticator(token))

		r.Post("/", createVote)

		r.Delete("/", deleteVote)
	})

	return r
}

// @Summary Создать новый голос
// @Description Создает новый голос на основе данных, переданных в теле запроса.
// @Tags Голоса
// @Accept json
// @Produce json
// @Param vote body services.VoteCreate true "Данные для создания голоса"
// @Success 201 {object} map[string]int "Успешно создано, возвращает идентификатор созданного голоса"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /votes [post]
func createVote(w http.ResponseWriter, r *http.Request) {
	id, err := RetrieveIdFromToken(r.Context())

	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	vc := services.VoteCreate{
		UserId: id,
	}

	if err := json.NewDecoder(r.Body).Decode(&vc); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = services.NewUsers().GetById(r.Context(), vc.UserId)

	if err != nil {
		http.Error(w, "Error finding user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = services.NewCandidates().GetById(r.Context(), vc.CandidateId)

	if err != nil {
		http.Error(w, "Error finding candidate: "+err.Error(), http.StatusInternalServerError)
		return
	}

	vote, err := services.NewVotes().Create(r.Context(), &vc)

	if err != nil {
		http.Error(w, "Error creating vote: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]int{"id": vote.ID})
}

// @Summary Удаление голоса пользователя
// @Description Удаляет голос пользователя по ID.
// @Tags Голоса
// @Accept json
// @Produce json
// @Param request body services.VoteDelete true "Информация о пользователе для удаления его голоса"
// @Success 204 {string} string "Голос успешно удален"
// @Failure 400 {string} string "Неверный формат данных или ошибка при удалении"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /votes [delete]
func deleteVote(w http.ResponseWriter, r *http.Request) {
	var vd services.VoteDelete

	if err := json.NewDecoder(r.Body).Decode(&vd); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := services.NewVotes().Delete(r.Context(), &vd)

	if err != nil {
		http.Error(w, "Error during vote deletion: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
