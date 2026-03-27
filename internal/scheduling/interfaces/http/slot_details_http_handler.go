package http

import (
	"fmt"
	"github/grzegab/calendar/internal/scheduling/application/schedule_details"
	"github/grzegab/calendar/internal/shared/infrastructure/auth"
	"net/http"
)

type SlotDetailsHttpHandler struct {
	svc *schedule_details.Handler
}

func NewSlotDetailsHttpHandler(svc *schedule_details.Handler) *SlotDetailsHttpHandler {
	return &SlotDetailsHttpHandler{svc: svc}
}

func (h *SlotDetailsHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	teacherID, ok := auth.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	errMsg := fmt.Sprintf("not implemented for teacher %s le", teacherID)
	http.Error(w, errMsg, http.StatusNotImplemented)
}
