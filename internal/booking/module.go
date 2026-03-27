package booking

import (
	"github/grzegab/calendar/internal/booking/application/cancel_booking"
	"github/grzegab/calendar/internal/booking/application/list_bookings"
	"github/grzegab/calendar/internal/booking/application/new_booking"
	"github/grzegab/calendar/internal/booking/domain"
	"github/grzegab/calendar/internal/booking/interfaces/http"
	"github/grzegab/calendar/internal/shared/infrastructure/auth"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Module struct {
	ListBookingsService      *list_bookings.Handler
	ListBookingsHttpHandler  *http.ListBookingHttpHandler
	NewBookingService        *new_booking.Handler
	NewBookingHttpHandler    *http.NewBookingHttpHandler
	CancelBookingService     *cancel_booking.Handler
	CancelBookingHttpHandler *http.CancelBookingHttpHandler
	tokenVerifier            auth.TokenVerifier
}

type Dependencies struct {
	DomainBookingRepository domain.BookingRepository
	ListBookingsRepository  list_bookings.ReadRepository
	TokenVerifier           auth.TokenVerifier
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
		tokenVerifier:            dep.TokenVerifier,
	}
}

func (m *Module) RegisterRoutes(r chi.Router) {
	r.Route("/bookings", func(r chi.Router) {
		r.Use(auth.JwtMiddleware(m.tokenVerifier))
		r.Use(middleware.Recoverer)

		r.Get("/", m.ListBookingsHttpHandler.Handle)
		r.Post("/", m.NewBookingHttpHandler.Handle)
		r.Delete("/{id}", m.CancelBookingHttpHandler.Handle)
	})
}
