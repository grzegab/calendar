package booking

import (
	"github/grzegab/calendar/internal/booking/domain"

	"github.com/go-chi/chi/v5"
)

type Module struct {
	ListBookingsService      *Service
	ListBookingsHttpHandler  *Handler
	NewBookingService        *Service
	NewBookingHttpHandler    *Handler
	CancelBookingService     *Service
	CancelBookingHttpHandler *Handler
}

type Dependencies struct {
	BookingRepository domain.BookingRepository
}

func NewModule(dep Dependencies) *Module {
	return &Module{
		Service: NewService(repo),
		Handler: NewHandler(repo),
	}
}

func (m *Module) RegisterRoutes(r chi.Router) {
	r.Route("/bookings", func(r chi.Router) {
		r.Get("/", m.ListBookingsHttpHandler.Handle)
		r.Post("/", m.NewBookingHttpHandler.Handle)
		r.Delete("/{id}", m.CancelBookingHttpHandler.Handle)
	})
}
