package main

import (
	"github.com/emirkosuta/celeritas"
	"github.com/emirkosuta/golang-laravel/data"
	"github.com/emirkosuta/golang-laravel/handlers"
	"github.com/emirkosuta/golang-laravel/middleware"
	"github.com/emirkosuta/golang-laravel/services"
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
