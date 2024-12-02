package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ElectionResultsRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getResults)

	return r
}

// TODO: implement logic
// TODO: create stored procedure to retrieve results
func getResults(w http.ResponseWriter, r *http.Request) {

}
