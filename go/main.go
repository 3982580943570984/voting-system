package main

import (
	"net/http"

	"voting-system/database"
	_ "voting-system/ent/generated/runtime"
	"voting-system/router"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	defer database.Client.Close()

	http.ListenAndServe(":3000", router.Router())
}
