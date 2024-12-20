package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Use(
		middleware.Heartbeat("/status"),
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler(cors.Options{
			AllowCredentials: true,
			AllowedHeaders:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			ExposedHeaders:   []string{"Link"},
			MaxAge:           300,
		}),
	)

	return router
}
