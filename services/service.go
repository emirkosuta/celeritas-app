package services

import (
	"myapp/data"

	"github.com/emirkosuta/celeritas"
)

type Services struct {
	App    *celeritas.Celeritas
	Models data.Models
}
