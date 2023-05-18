package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"myapp/services"

	"github.com/emirkosuta/celeritas"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Services   *services.Services
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
