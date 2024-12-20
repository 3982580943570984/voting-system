package routes

import (
	"elections/repositories"
	"encoding/json"
	"net/http"
	"shared/token"
	"shared/utils"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func Comments() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token.Token), jwtauth.Authenticator(token.Token))

		r.Post("/", createComment)

		r.Get("/", getAllCommentsByElectionId)
	})

	return r
}

// @Summary Создать комментарий
// @Description Создает новый комментарий для указанного выбора.
// @Tags Комментарии
// @Accept  json
// @Produce  json
// @Param id path int true "ID выборов"
// @Param comment body services.CommentCreate true "Данные для создания комментария"
// @Success 201 {object} map[string]int "Успешно создано, возвращает идентификатор созданного комментария"
// @Failure 400 {string} string "Неверный формат запроса или ID выборов"
// @Failure 401 {string} string "Неавторизованный доступ"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id}/comments [post]]
// @Security Bearer
func createComment(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.RetrieveIdFromToken(r.Context())
	if err != nil {
		http.Error(w, "Invalid user ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	electionId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	cc := repositories.CommentCreate{UserId: userId, ElectionId: electionId}
	if err := json.NewDecoder(r.Body).Decode(&cc); err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	comment, err := repositories.NewComments().Create(r.Context(), &cc)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]int{"id": comment.ID})
}

// @Summary Получить все комментарии для выборов
// @Description Возвращает список всех комментариев, связанных с указанными выборами.
// @Tags Комментарии
// @Accept json
// @Produce json
// @Param id path int true "ID выборов"
// @Success 200 {array} generated.Comment "Список комментариев"
// @Failure 400 {string} string "Неверный формат ID выборов"
// @Failure 401 {string} string "Неавторизованный доступ"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /elections/{id}/comments [get]
// @Security Bearer
func getAllCommentsByElectionId(w http.ResponseWriter, r *http.Request) {
	electionId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	comments, err := repositories.NewComments().GetAllByElectionId(r.Context(), electionId)
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(comments)
}
