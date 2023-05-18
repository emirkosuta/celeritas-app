package main

import (
	"log"
	"os"

	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"myapp/services"

	"github.com/emirkosuta/celeritas"
)

func initApplication() *application {
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

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	myHandlers := &handlers.Handlers{
		App: cel,
	}

	myServices := &services.Services{
		App: cel,
	}

	app := &application{
		App:        cel,
		Handlers:   myHandlers,
		Services:   myServices,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Services = app.Services
	myServices.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
