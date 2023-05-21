package main

import (
	"myapp/middleware"

	"github.com/emirkosuta/celeritas"
)

type application struct {
	App        *celeritas.Celeritas
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
