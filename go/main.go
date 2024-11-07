package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"voting-system/authentication"
	"voting-system/users"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

// TODO: Move `create<table_name>Table` to storage files in packages

var tokenAuth *jwtauth.JWTAuth

var usersResource *users.Resource

func init() {
	// TODO: replace with secret key
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func main() {
	database := openDatabase()

	defer database.Close()

	createTables(database)

	usersResource = users.NewResource(users.NewStorage(database))

	http.ListenAndServe(":3000", router())
}

func router() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var credentials authentication.Credentials

		error := json.NewDecoder(r.Body).Decode(&credentials)

		if error != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		user, error := usersResource.Storage.GetByEmail(credentials.Email)

		if error != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		error = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))

		if error != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		_, tokenString, error := tokenAuth.Encode(map[string]interface{}{"user_id": user.Id})

		if error != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	})

	r.Post("/signup", func(w http.ResponseWriter, r *http.Request) {
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

		user, error := usersResource.Storage.GetByEmail(credentials.Email)

		if error != nil && error != sql.ErrNoRows {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if user != nil {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		hashedPassword, error := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)

		if error != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		newUser := users.User{
			Email:    credentials.Email,
			Password: string(hashedPassword),
			IsActive: true,
		}

		userID, error := usersResource.Storage.Create(newUser)

		if error != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User created successfully",
			"user_id": userID,
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))

		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())

			userID := int(claims["user_id"].(float64))

			user, err := usersResource.Storage.GetByID(userID)

			if err != nil {
				http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
				return
			}

			if user == nil {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}

			w.Write([]byte(fmt.Sprintf("Protected area. Hi %s", user.Email)))
		})
	})

	// TODO: debug only endpoints. remove later
	r.Mount("/users", usersResource.Routes())

	return r
}

// TODO: move to database folder

func openDatabase() *sql.DB {
	connectionString := "user=postgres dbname=database host=/run/postgresql sslmode=disable"

	database, error := sql.Open("postgres", connectionString)

	if error != nil {
		log.Fatalf("Failed to open database: %v", error)
	}

	return database
}

func createTable(db *sql.DB, query string) {
	_, error := db.Exec(query)

	if error != nil {
		log.Fatalf("Failed to create table: %v", error)
	}
}

func createUsersTable(db *sql.DB) {
	createTable(db, `
CREATE TABLE IF NOT EXISTS users (
	user_id SERIAL PRIMARY KEY,
	email VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	is_active BOOLEAN DEFAULT TRUE
);
`)
}

func createProfilesTable(db *sql.DB) {
	createTable(db, `
CREATE TABLE IF NOT EXISTS profiles (
	user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	birthdate DATE
);
`)
}

func createVotersTable(db *sql.DB) {
	createTable(db, `
CREATE TABLE IF NOT EXISTS voters (
	voter_id SERIAL PRIMARY KEY,
	user_id INT REFERENCES users(user_id) ON DELETE CASCADE
);
`)
}

func createOrganizersTable(db *sql.DB) {
	createTable(db, `
CREATE TABLE IF NOT EXISTS organizers (
	organizer_id SERIAL PRIMARY KEY,
	user_id INT REFERENCES users(user_id) ON DELETE CASCADE
);
`)
}

func createElectionsTable(db *sql.DB) {
	createTable(db, `
CREATE TABLE IF NOT EXISTS elections (
	election_id SERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	start_date DATE,
	end_date DATE,
	is_active BOOLEAN DEFAULT TRUE
);
`)
}

func createCandidatesTable(db *sql.DB) {
	createTable(db, `
CREATE TABLE IF NOT EXISTS candidates (
	candidate_id SERIAL PRIMARY KEY,
	election_id INT REFERENCES elections(election_id) ON DELETE CASCADE,
	name VARCHAR(255) NOT NULL,
	description TEXT
);
`)
}

func createVotesTable(db *sql.DB) {
	createTable(db, `
CREATE TABLE IF NOT EXISTS votes (
	vote_id SERIAL PRIMARY KEY,
	user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
	election_id INT REFERENCES elections(election_id) ON DELETE CASCADE,
	candidate_id INT REFERENCES candidates(candidate_id) ON DELETE CASCADE,
	timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	is_active BOOLEAN DEFAULT TRUE
);
`)
}

func createTables(db *sql.DB) {
	createUsersTable(db)

	createProfilesTable(db)

	createVotersTable(db)

	createOrganizersTable(db)

	createElectionsTable(db)

	createCandidatesTable(db)

	createVotesTable(db)
}
