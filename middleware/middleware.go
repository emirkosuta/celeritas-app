package middleware

import (
	"myapp/data"

	"github.com/emirkosuta/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
