package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"voting-system/services"

	"github.com/go-chi/chi/v5"
)

func UsersRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getAllUsers)

	r.Get("/{id}", getUser)

	return r
}

// @Summary		Получить пользователя
// @Description	Получает пользователя по уникальному идентификатору
// @Tags			Пользователи
// @Accept			json
// @Produce		json
// @Param			id	path		int					true	"ID пользователя"
// @Success		200	{object}	generated.User		"Данные пользователя"
// @Failure		400	{object}	map[string]string	"Неверный ID"
// @Failure		404	{object}	map[string]string	"Пользователь не найден"
// @Router			/users/{id} [get]
func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := services.NewUsers().
		GetById(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

// @Summary		Получить всех пользователей
// @Description	Получает список всех зарегистрированных пользователей
// @Tags			Пользователи
// @Accept			json
// @Produce		json
// @Success		200	{array}		generated.User		"Список пользователей"
// @Failure		500	{object}	map[string]string	"Внутренняя ошибка сервера"
// @Router			/users [get]
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.NewUsers().
		GetAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}
