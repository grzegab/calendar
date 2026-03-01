package http

import (
	"github/grzegab/calendar/internal/booking/application/list_bookings"
	"net/http"
)

type ListBookingHttpHandler struct {
	service *list_bookings.Handler
}

func NewListBookingHttpHandler(svc *list_bookings.Handler) *ListBookingHttpHandler {
	return &ListBookingHttpHandler{service: svc}
}

func (h *ListBookingHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {}
