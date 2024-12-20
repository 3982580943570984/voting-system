package main

import (
	"elections/router"
	"elections/scheduler"
	"log"
	"net/http"
	"shared/database"

	_ "shared/ent/generated/runtime"
)

func main() {
	defer database.Client.Close()

	defer scheduler.Scheduler.Shutdown()

	log.Print("Started")

	http.ListenAndServe(":3002", router.Router())
}
