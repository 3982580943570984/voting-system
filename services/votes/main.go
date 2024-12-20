package main

import (
	"log"
	"net/http"
	"shared/database"
	"votes/router"

	_ "shared/ent/generated/runtime"
)

func main() {
	defer database.Client.Close()

	log.Print("Started")

	http.ListenAndServe(":3003", router.Router())
}
