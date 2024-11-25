package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"voting-system/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

// Credentials представляет учетные данные пользователя для входа/регистрации.
// swagger:model Credentials
type Credentials struct {
	// Электронная почта пользователя
	// required: true
	Email string `json:"email"`

	// Пароль пользователя
	// required: true
	Password string `json:"password"`
}

// TokenResponse представляет ответ с JWT-токеном.
// swagger:model TokenResponse
type TokenResponse struct {
	// JWT-токен
	Token string `json:"token"`
}

// SignupResponse представляет ответ после успешной регистрации.
// swagger:model SignupResponse
type SignupResponse struct {
	// ID вновь созданного пользователя
	ID int `json:"id"`

	// Сообщение об успешном создании пользователя
	Message string `json:"message"`
}

var token *jwtauth.JWTAuth

func init() {
	// TODO: replace with secret key
	token = jwtauth.New("HS256", []byte("secret"), nil)
}

func AuthenticationRoutes() chi.Router {
	router := chi.NewRouter()

	router.Post("/login", login)

	router.Post("/signup", signup)

	// TODO: remove. was needed for debug purposes
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token))

		r.Use(jwtauth.Authenticator(token))

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, err := jwtauth.FromContext(r.Context())

			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			idFloat, ok := claims["id"].(float64)

			if !ok {
				http.Error(w, "Invalid or missing 'id' in token", http.StatusBadRequest)
				return
			}

			id := int(idFloat)

			user, err := services.
				NewUsers().
				GetById(r.Context(), id)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			w.Write([]byte(fmt.Sprintf("Protected area. Hi %s", user.Email)))
		})
	})

	return router
}

// @Summary		Вход пользователя
// @Description	Аутентификация пользователя и возврат JWT-токена
// @Tags			Аутентификация
// @Accept			json
// @Produce		json
// @Param			credentials	body		Credentials	true	"Учетные данные пользователя"
// @Success		200			{object}	TokenResponse
// @Failure		400			{object}	map[string]string	"Неверный формат запроса"
// @Failure		401			{object}	map[string]string	"Неверный пароль"
// @Failure		500			{object}	map[string]string	"Внутренняя ошибка сервера"
// @Router			/login [post]
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	_, tokenString, err := token.Encode(map[string]interface{}{"id": user.ID})

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	loginTime := time.Now()
	_, err = service.Update(r.Context(), &services.UserUpdate{
		ID:        user.ID,
		LastLogin: &loginTime,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(TokenResponse{Token: tokenString})
}

// @Summary		Регистрация пользователя
// @Description	Регистрация нового пользователя
// @Tags			Аутентификация
// @Accept			json
// @Produce		json
// @Param			credentials	body		Credentials	true	"Учетные данные пользователя"
// @Success		201			{object}	SignupResponse
// @Failure		400			{object}	map[string]string	"Неверный формат запроса или ошибка при создании пользователя"
// @Router			/signup [post]
func signup(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := services.
		NewUsers().
		Create(r.Context(), &services.UserCreate{
			Email:    creds.Email,
			Password: creds.Password,
		})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(SignupResponse{
		Message: "Пользователь успешно создан",
		ID:      user.ID,
	})
}
