package main

import (
	"log"
	"os"

	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/emirkosuta/celeritas"
)

func initApplication() *celeritas.Celeritas {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"

	models := data.New(cel.DB.Pool)

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	myHandlers := &handlers.Handlers{}

	cel.Routes = routes(cel, myMiddleware, myHandlers)

	return cel
}
