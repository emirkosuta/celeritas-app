package middleware

import (
	"github.com/emirkosuta/celeritas"
	"github.com/emirkosuta/golang-laravel/data"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
