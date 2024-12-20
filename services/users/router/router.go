package router

import (
	"encoding/json"
	"net/http"
	"strconv"
	"users/repositories"

	"shared/token"
	"shared/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Use(
		middleware.Heartbeat("/status"),
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler(cors.Options{
			AllowCredentials: true,
			AllowedHeaders:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			ExposedHeaders:   []string{"Link"},
			MaxAge:           300,
		}),
	)

	router.Post("/login", login)

	router.Post("/signup", signup)

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token.Token), jwtauth.Authenticator(token.Token))

		r.Get("/profile", getUserProfile)
	})

	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", get)

		r.Get("/profile", getProfile)
	})

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
	credentials := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	repository := repositories.NewUsers()

	user, err := repository.GetByEmail(r.Context(), credentials.Email)
	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	_, tokenString, err := token.Token.Encode(map[string]interface{}{"id": user.ID, "is_organizer": user.IsOrganizer})
	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = repository.Update(r.Context(), &repositories.UserUpdate{ID: user.ID})
	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tokenString)
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
	credentials := struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		IsOrganizer bool   `json:"is_organizer"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	uc := repositories.UserCreate{
		Email:       credentials.Email,
		Password:    credentials.Password,
		IsOrganizer: credentials.IsOrganizer,
	}

	user, err := repositories.NewUsers().Create(r.Context(), &uc)
	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]int{"id": user.ID})
}

func get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID"+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := repositories.NewUsers().GetById(r.Context(), id)
	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func getUserProfile(w http.ResponseWriter, r *http.Request) {
	id, err := utils.RetrieveIdFromToken(r.Context())
	if err != nil {
		http.Error(w, "Invalid user ID"+err.Error(), http.StatusBadRequest)
		return
	}

	profile, err := repositories.NewProfiles().GetByUserId(r.Context(), id)
	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(profile)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID"+err.Error(), http.StatusBadRequest)
		return
	}

	profile, err := repositories.NewProfiles().GetByUserId(r.Context(), id)
	if err != nil {
		http.Error(w, "Internal server error"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(profile)
}
