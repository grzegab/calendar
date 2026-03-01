package booking

import (
	"github/grzegab/calendar/internal/booking/application/cancel_booking"
	"github/grzegab/calendar/internal/booking/application/list_bookings"
	"github/grzegab/calendar/internal/booking/application/new_booking"
	"github/grzegab/calendar/internal/booking/domain"
	"github/grzegab/calendar/internal/booking/interfaces/http"

	"github.com/go-chi/chi/v5"
)

type Module struct {
	ListBookingsService      *list_bookings.Handler
	ListBookingsHttpHandler  *http.ListBookingHttpHandler
	NewBookingService        *new_booking.Handler
	NewBookingHttpHandler    *http.NewBookingHttpHandler
	CancelBookingService     *cancel_booking.Handler
	CancelBookingHttpHandler *http.CancelBookingHttpHandler
}

type Dependencies struct {
	DomainBookingRepository domain.BookingRepository
	ListBookingsRepository  list_bookings.ReadRepository
}

func NewModule(dep Dependencies) *Module {
	listBookingsService := list_bookings.NewHandler(dep.ListBookingsRepository)
	listBookingsHttpHandler := http.NewListBookingHttpHandler(listBookingsService)
	newBookingService := new_booking.NewHandler(dep.DomainBookingRepository)
	newBookingHttpHandler := http.NewNewBookingHttpHandler(newBookingService)
	cancelBookingService := cancel_booking.NewHandler(dep.DomainBookingRepository)
	cancelBookingHttpHandler := http.NewCancelBookingHandler(cancelBookingService)

	return &Module{
		ListBookingsService:      listBookingsService,
		ListBookingsHttpHandler:  listBookingsHttpHandler,
		NewBookingService:        newBookingService,
		NewBookingHttpHandler:    newBookingHttpHandler,
		CancelBookingService:     cancelBookingService,
		CancelBookingHttpHandler: cancelBookingHttpHandler,
	}
}

func (m *Module) RegisterRoutes(r chi.Router) {
	r.Route("/bookings", func(r chi.Router) {
		r.Get("/", m.ListBookingsHttpHandler.Handle)
		r.Post("/", m.NewBookingHttpHandler.Handle)
		r.Delete("/{id}", m.CancelBookingHttpHandler.Handle)
	})
}
