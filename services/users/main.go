package main

import (
	"log"
	"net/http"
	"shared/database"
	"users/router"

	_ "shared/ent/generated/runtime"
)

func main() {
	defer database.Client.Close()

	log.Print("Started")

	http.ListenAndServe(":3001", router.Router())
}
