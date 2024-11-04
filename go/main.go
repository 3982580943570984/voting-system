package main

import (
	"database/sql"
	"log"
	"net/http"

	"voting-system/users"

	"github.com/go-chi/chi/v5"

	_ "github.com/lib/pq"
)

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

// TODO: Move `create<table_name>Table` to storage files in packages

func main() {
	database := openDatabase()

	defer database.Close()

	createTables(database)

	usersStorage := users.NewStorage(database)

	usersResource := users.NewResource(usersStorage)

	r := chi.NewRouter()

	r.Mount("/users", usersResource.Routes())

	http.ListenAndServe(":3000", r)
}
