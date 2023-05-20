package main

import (
	"myapp/handlers"
	"myapp/middleware"

	"github.com/emirkosuta/celeritas"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
