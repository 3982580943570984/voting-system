package routes

import (
	"encoding/json"
	"net/http"
	"voting-system/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

var token *jwtauth.JWTAuth

func init() {
	// TODO: replace with secret key
	token = jwtauth.New("HS256", []byte("secret"), nil)
}

// Credentials представляет учетные данные пользователя для входа/регистрации.
// swagger:model Credentials
type Credentials struct {
	// Электронная почта пользователя
	// required: true
	Email string `json:"email"`

	// Пароль пользователя
	// required: true
	Password string `json:"password"`

	// Пользователь является организатором.
	IsOrganizer bool `json:"is_organizer"`
}

// TokenResponse представляет ответ с JWT-токеном.
// swagger:model TokenResponse
type TokenResponse struct {
	// JWT-токен
	Token string `json:"token"`
}

func AuthenticationRoutes() chi.Router {
	router := chi.NewRouter()

	router.Post("/login", login)

	router.Post("/signup", signup)

	return router
}

// @Summary Вход пользователя
// @Description Аутентификация пользователя и возврат JWT-токена
// @Tags Аутентификация
// @Accept json
// @Produce json
// @Param credentials body Credentials true "Учетные данные пользователя"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} map[string]string "Неверный формат запроса"
// @Failure 401 {object} map[string]string "Неверный пароль"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /login [post]
func login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	service := services.NewUsers()

	user, err := service.
		GetByEmail(r.Context(), creds.Email)

	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	_, tokenString, err := token.Encode(map[string]interface{}{"id": user.ID, "is_organizer": user.IsOrganizer})

	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = service.Update(r.Context(), &services.UserUpdate{ID: user.ID})

	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(TokenResponse{Token: tokenString})
}

// @Summary Регистрация пользователя
// @Description Регистрация нового пользователя
// @Tags Аутентификация
// @Accept json
// @Produce json
// @Param credentials body Credentials true "Учетные данные пользователя"
// @Success 201 {object} int "Успешно создано, возвращает идентификатор созданного пользователя"
// @Failure 400 {object} map[string]string "Неверный формат запроса или ошибка при создании пользователя"
// @Router /signup [post]
func signup(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	uc := services.UserCreate{
		Email:       creds.Email,
		Password:    creds.Password,
		IsOrganizer: creds.IsOrganizer,
	}

	user, err := services.NewUsers().Create(r.Context(), &uc)

	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]int{"id": user.ID})
}
