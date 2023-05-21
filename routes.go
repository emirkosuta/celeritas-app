package main

import (
	"myapp/handlers"
	"myapp/middleware"

	"github.com/emirkosuta/celeritas"
	"github.com/go-chi/chi/v5"
)

func routes(app *celeritas.Celeritas, middleware *middleware.Middleware, handlers *handlers.Handlers) *chi.Mux {
	// middleware must come before any routes

	// add routes here
	app.Routes.Get("/", handlers.Home)

	//api
	app.Routes.Route("/api", func(rt chi.Router) {
		rt.Post("/cars", handlers.CarHandler.CreateCar)
		rt.Get("/cars/{id}", handlers.CarHandler.GetCarById)
		rt.Get("/cars", handlers.CarHandler.GetCarList)
		rt.Put("/cars/{id}", handlers.CarHandler.UpdateCar)
		rt.Delete("/cars/{id}", handlers.CarHandler.DeleteCar)
	})

	return app.Routes
}
