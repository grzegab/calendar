package http

import (
	"github/grzegab/calendar/internal/booking/application/cancel_booking"
	"net/http"
)

type CancelBookingHttpHandler struct {
	serivce *cancel_booking.Handler
}

func NewCancelBookingHandler(svc *cancel_booking.Handler) *CancelBookingHttpHandler {
	return &CancelBookingHttpHandler{serivce: svc}
}

func (h *CancelBookingHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {}
