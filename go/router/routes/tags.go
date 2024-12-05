package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"voting-system/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func TagsRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getElectionTags)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token), jwtauth.Authenticator(token))

		r.Post("/", createElectionTag)

		r.Put("/", updateElectionTags)
	})

	return r
}

// @Summary Создать тег для выборов
// @Description Создает новый тег и ассоциирует его с указанными выборами.
// @Tags Теги
// @Accept  json
// @Produce  json
// @Param id path int true "ID выборов"
// @Param tag body services.TagCreate true "Данные для создания тега (только имя)"
// @Success 201 {object} services.TagCreate "ID созданного тега"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /elections/{id}/tags [post]
// @Security Bearer
func createElectionTag(w http.ResponseWriter, r *http.Request) {
	electionId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid election ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	tc := services.TagCreate{ElectionId: electionId}
	if err := json.NewDecoder(r.Body).Decode(&tc); err != nil {
		http.Error(w, "Неверное тело запроса: "+err.Error(), http.StatusBadRequest)
		return
	}

	tag, err := services.NewTags().Create(r.Context(), &tc)
	if err != nil {
		http.Error(w, "Error creating tag: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]int{"id": tag.ID})
}

// @Summary Получить теги выборов
// @Description Возвращает список тегов, связанных с конкретными выборами.
// @Tags Теги
// @Accept  json
// @Produce  json
// @Param id path int true "ID выборов"
// @Success 200 {array} generated.Tag "Список тегов"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /elections/{id}/tags [get]
func getElectionTags(w http.ResponseWriter, r *http.Request) {
	electionId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid election ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	tags, err := services.NewTags().GetByElectionId(r.Context(), electionId)
	if err != nil {
		http.Error(w, "Error retrieving tags: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tags)
}

// @Summary Обновить теги выборов
// @Description Обновляет теги, связанные с конкретными выборами. Все существующие теги будут заменены на предоставленные.
// @Tags Теги
// @Accept  json
// @Produce  json
// @Param id path int true "ID выборов"
// @Param tags body services.TagsUpdate true "Новый список тегов"
// @Success 204 {string} string "Теги успешно обновлены"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /elections/{id}/tags [put]
// @Security Bearer
func updateElectionTags(w http.ResponseWriter, r *http.Request) {
	electionId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid election ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	tu := services.TagsUpdate{ElectionId: electionId}
	if err := json.NewDecoder(r.Body).Decode(&tu); err != nil {
		http.Error(w, "Неверное тело запроса: "+err.Error(), http.StatusBadRequest)
		return
	}

	if tu.Names == nil {
		http.Error(w, "Поле 'names' обязательно для заполнения", http.StatusBadRequest)
		return
	}

	err = services.NewTags().UpdateByElectionId(r.Context(), &tu)
	if err != nil {
		http.Error(w, "Ошибка при обновлении тегов: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
