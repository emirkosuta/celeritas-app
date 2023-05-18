package main

import (
	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// add routes here
	a.App.Routes.Get("/", a.Handlers.Home)

	//api
	a.App.Routes.Route("/api", func(rt chi.Router) {

	})

	return a.App.Routes
}
