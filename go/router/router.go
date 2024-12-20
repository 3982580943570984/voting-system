package router

import (
	"net/http"
	"voting-system/router/routes"

	_ "voting-system/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	router.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "swagger.json")
	})

	router.Mount("/elections", routes.ElectionsRoutes())

	router.Mount("/", routes.AuthenticationRoutes())

	router.Mount("/votes", routes.VotesRoutes())

	router.Mount("/profiles", routes.ProfileRoutes())

	return router
}
