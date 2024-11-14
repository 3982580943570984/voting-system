package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"voting-system/database"

	"github.com/go-chi/chi/v5"
)

type userInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
}

func UsersRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getAllUsers)
	r.Post("/", createUser)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", getUser)
		r.Put("/", updateUser)
		r.Delete("/", deleteUser)
	})

	return r
}

// createUser создает нового пользователя
//
//	@Summary		Создать пользователя
//	@Description	Создает нового пользователя с указанными данными
//	@Tags			Пользователи
//	@Accept			json
//	@Produce		json
//	@Param			user	body		userInput			true	"Данные пользователя"
//	@Success		200		{object}	map[string]int		"ID созданного пользователя"
//	@Failure		400		{object}	map[string]string	"Неверный запрос"
//	@Router			/users [post]
func createUser(w http.ResponseWriter, r *http.Request) {
	var input userInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user := database.Users.
		Create().
		SetEmail(input.Email).
		SetPassword(input.Password).
		SetIsActive(input.IsActive).
		SaveX(r.Context())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]int{"id": user.ID})
}

// getUser получает пользователя по ID
// @Summary Получить пользователя
// @Description Получает пользователя по уникальному идентификатору
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} User "Данные пользователя"
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Router /users/{id} [get]
func getUser(w http.ResponseWriter, r *http.Request) {
	id, error := strconv.Atoi(chi.URLParam(r, "id"))
	if error != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user := database.Users.GetX(r.Context(), id)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

// getAllUsers получает список всех пользователей
// @Summary Получить всех пользователей
// @Description Получает список всех зарегистрированных пользователей
// @Tags Пользователи
// @Accept json
// @Produce json
// @Success 200 {array} User "Список пользователей"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /users [get]
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := database.Users.
		Query().
		AllX(r.Context())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

// updateUser обновляет данные пользователя по ID
// @Summary Обновить пользователя
// @Description Обновляет данные существующего пользователя по уникальному идентификатору
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param user body userInput true "Обновленные данные пользователя"
// @Success 200 {object} map[string]int "ID обновленного пользователя"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Router /users/{id} [put]
func updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var input userInput
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedUser := database.Users.
		UpdateOneID(id).
		SetEmail(input.Email).
		SetPassword(input.Password).
		SetIsActive(input.IsActive).
		SaveX(r.Context())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]int{"id": updatedUser.ID})
}

// deleteUser удаляет пользователя по ID
// @Summary Удалить пользователя
// @Description Удаляет существующего пользователя по уникальному идентификатору
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} map[string]int "ID удаленного пользователя"
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Router /users/{id} [delete]
func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	database.Users.
		DeleteOneID(id).
		ExecX(r.Context())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]int{"id": id})
}
