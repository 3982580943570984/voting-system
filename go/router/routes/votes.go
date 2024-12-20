package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"voting-system/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func VotesRoutes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token), jwtauth.Authenticator(token))

		r.Get("/", getVotes)

		r.Post("/", createVote)

		r.Get("/voted", hasVoted)
	})

	return r
}

// @Summary Получить голоса пользователя на конкретных выборах
// @Description Возвращает список голосов текущего пользователя на указанных выборах.
// @Tags Голоса
// @Security Bearer
// @Accept json
// @Produce json
// @Param election_id query int true "ID выборов"
// @Success 200 {array} generated.Vote "Список голосов пользователя"
// @Failure 400 {object} map[string]string "Неверный формат запроса или отсутствует параметр"
// @Failure 401 {object} map[string]string "Неавторизованный доступ"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /votes [get]
func getVotes(w http.ResponseWriter, r *http.Request) {
	id, err := RetrieveIdFromToken(r.Context())

	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	electionIdStr := r.URL.Query().Get("election_id")
	if electionIdStr == "" {
		http.Error(w, "Отсутствует параметр election_id", http.StatusBadRequest)
		return
	}

	electionId, err := strconv.Atoi(electionIdStr)
	if err != nil {
		http.Error(w, "Неверный параметр election_id: "+err.Error(), http.StatusBadRequest)
		return
	}

	votes, err := services.NewVotes().GetByUserAndElectionId(r.Context(), id, electionId)
	if err != nil {
		http.Error(w, "Ошибка при получении голосов: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(votes)
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

// @Summary Проверить, проголосовал ли пользователь на выборах
// @Description Возвращает булевое значение, указывающее, проголосовал ли пользователь на указанных выборах.
// @Tags Голоса
// @Security Bearer
// @Accept json
// @Produce json
// @Param election_id query int true "ID выборов"
// @Success 200 {object} map[string]bool "Статус голосования пользователя"
// @Failure 400 {object} map[string]string "Неверный формат запроса или отсутствует параметр"
// @Failure 401 {object} map[string]string "Неавторизованный доступ"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /votes/voted [get]
func hasVoted(w http.ResponseWriter, r *http.Request) {
	userId, err := RetrieveIdFromToken(r.Context())

	if err != nil {
		http.Error(w, "Неверные данные: "+err.Error(), http.StatusBadRequest)
		return
	}

	electionIdStr := r.URL.Query().Get("election_id")
	if electionIdStr == "" {
		http.Error(w, "Отсутствует параметр election_id", http.StatusBadRequest)
		return
	}

	electionId, err := strconv.Atoi(electionIdStr)
	if err != nil {
		http.Error(w, "Неверный параметр election_id: "+err.Error(), http.StatusBadRequest)
		return
	}

	votes, err := services.NewVotes().GetByUserAndElectionId(r.Context(), userId, electionId)
	if err != nil {
		http.Error(w, "Ошибка при получении голосов: "+err.Error(), http.StatusInternalServerError)
		return
	}

	hasVoted := len(votes) > 0

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"voted": hasVoted})
}
