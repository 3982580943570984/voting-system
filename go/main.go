package main

import (
	"net/http"

	"voting-system/database"
	"voting-system/router"
)

// TODO: change all database operations to panic
// TODO: write open api file
// TODO: change entity names
// TODO: populate entity attributes

// TODO: All fields are required by default, and can be set to optional using the Optional method.
// TODO: consider to use pgx instead of pq

func main() {
	// TODO: consider to remove or smth similar
	defer database.Client.Close()

	http.ListenAndServe(":3000", router.Router())
}
