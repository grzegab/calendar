package http

import (
	"github/grzegab/calendar/internal/booking/application/new_booking"
	"net/http"
)

type NewBookingHttpHandler struct {
	service *new_booking.Handler
}

func NewNewBookingHttpHandler(svc *new_booking.Handler) *NewBookingHttpHandler {
	return &NewBookingHttpHandler{service: svc}
}

func (h *NewBookingHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {}
