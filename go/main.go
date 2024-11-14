package main

import (
	"net/http"

	"voting-system/database"
	"voting-system/router"
)

// TODO: change all database operations to panic
// TODO: write open api file

func main() {
	// TODO: consider to remove or smth similar
	defer database.Client.Close()

	http.ListenAndServe(":3000", router.Router())
}
