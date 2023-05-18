package services

import (
	"github.com/emirkosuta/celeritas"
	"github.com/emirkosuta/golang-laravel/data"
)

type Services struct {
	App    *celeritas.Celeritas
	Models data.Models
}
