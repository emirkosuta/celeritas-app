package handlers

import (
	"net/http"
	"time"

	"github.com/emirkosuta/celeritas"
	"github.com/emirkosuta/golang-laravel/services"
)

type Handlers struct {
	App      *celeritas.Celeritas
	Services *services.Services
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
