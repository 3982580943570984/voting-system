package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"voting-system/authentication"
	"voting-system/database"
	"voting-system/ent"
	"voting-system/ent/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

func AuthenticationRoutes() chi.Router {
	router := chi.NewRouter()

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var credentials authentication.Credentials

		err := json.NewDecoder(r.Body).Decode(&credentials)

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		user, err := database.Users.
			Query().
			Where(user.EmailEQ(credentials.Email)).
			Only(context.Background())

		if err != nil {
			if ent.IsNotFound(err) {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))

		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		_, tokenString, err := authentication.Token.Encode(map[string]interface{}{"id": user.ID})

		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	})

	router.Post("/signup", func(w http.ResponseWriter, r *http.Request) {
		var credentials authentication.Credentials

		error := json.NewDecoder(r.Body).Decode(&credentials)

		if error != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if credentials.Email == "" || credentials.Password == "" {
			http.Error(w, "Email and password are required", http.StatusBadRequest)
			return
		}

		userExists, error := database.Users.
			Query().
			Where(user.EmailEQ(credentials.Email)).
			Exist(context.Background())

		if error != nil {
			log.Printf("Internal server error: %v", error)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if userExists {
			log.Print("User already exists")
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		hashedPassword, error := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)

		if error != nil {
			log.Printf("Internal server error: %v", error)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		newUser := ent.User{
			Email:    credentials.Email,
			Password: string(hashedPassword),
			IsActive: true,
		}

		createdUser, error := database.Users.
			Create().
			SetEmail(newUser.Email).
			SetPassword(newUser.Password).
			SetIsActive(newUser.IsActive).
			Save(context.Background())

		if error != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User created successfully",
			"id":      createdUser.ID,
		})
	})

	// TODO: remove. was needed for debug purposes
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(authentication.Token))

		r.Use(jwtauth.Authenticator(authentication.Token))

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

			user, err := database.Users.Get(r.Context(), id)

			if err != nil {
				if user == nil {
					http.Error(w, "User not found", http.StatusNotFound)
					return
				}

				http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
				return
			}

			w.Write([]byte(fmt.Sprintf("Protected area. Hi %s", user.Email)))
		})
	})

	return router
}
