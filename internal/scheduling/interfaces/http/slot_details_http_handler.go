package http

import (
	"github/grzegab/calendar/internal/scheduling/application/schedule_details"
	"net/http"
)

type SlotDetailsHttpHandler struct {
	svc *schedule_details.Handler
}

func NewSlotDetailsHttpHandler(svc *schedule_details.Handler) *SlotDetailsHttpHandler {
	return &SlotDetailsHttpHandler{svc: svc}
}

func (h *SlotDetailsHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {

}
