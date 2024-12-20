package main

import (
	"comments/router"
	"log"
	"net/http"
	"shared/database"

	_ "shared/ent/generated/runtime"
)

func main() {
	defer database.Client.Close()

	log.Print("Started")

	http.ListenAndServe(":3004", router.Router())
}
