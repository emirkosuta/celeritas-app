package main

import (
	"myapp/handlers"
	"myapp/middleware"

	"github.com/emirkosuta/celeritas"
	"github.com/go-chi/chi/v5"
)

func routes(app *celeritas.Celeritas, middleware *middleware.Middleware, handlers *handlers.Handlers) *chi.Mux {
	// middleware must come before any routes

	//api
	app.Routes.Route("/api", func(rt chi.Router) {

	})

	return app.Routes
}
