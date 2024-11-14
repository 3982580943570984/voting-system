package router

import (
	"encoding/json"
	"net/http"
	"voting-system/router/routes"

	_ "voting-system/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Use(middleware.Recoverer)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Mount("/", routes.AuthenticationRoutes())

	router.Mount("/users", routes.UsersRoutes())

	router.Mount("/elections", routes.ElectionsRoutes())

	router.Mount("/profiles", routes.ProfilesRoutes())

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("swagger/doc.json"),
	))

	return router
}
