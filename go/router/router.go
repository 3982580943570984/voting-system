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

	router.Use(middleware.Logger, middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Mount("/", routes.AuthenticationRoutes())

	router.Mount("/elections", routes.ElectionsRoutes())

	router.Mount("/votes", routes.VotesRoutes())

	router.Mount("/profiles", routes.ProfileRoutes())

	router.Mount("/tags", routes.TagsRoutes())

	router.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "swagger.json")
	})

	return router
}
