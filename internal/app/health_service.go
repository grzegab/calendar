package app

import (
	"github/grzegab/calendar/internal/app/http"

	"github.com/go-chi/chi/v5"
)

type HealthService struct {
	pingHandler http.PingHandler
}

func NewHealthService() *HealthService {
	return &HealthService{
		pingHandler: http.PingHandler{},
	}
}

func (h *HealthService) RegisterRoutes(r chi.Router) {
	r.Route("/health", func(r chi.Router) {
		r.Get("/ping", h.pingHandler.Handle)
	})
}
