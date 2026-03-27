package http

import (
	"fmt"
	"github/grzegab/calendar/internal/scheduling/application/confirm_schedule"
	"github/grzegab/calendar/internal/shared/infrastructure/auth"
	"net/http"
)

type ConfirmScheduleHttpHandler struct {
	service *confirm_schedule.Handler
}

func CreateConfirmScheduleHttpHandler(svc *confirm_schedule.Handler) *ConfirmScheduleHttpHandler {
	return &ConfirmScheduleHttpHandler{service: svc}
}

func (h *ConfirmScheduleHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	teacherID, ok := auth.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	errMsg := fmt.Sprintf("not implemented for teacher %s le", teacherID)
	http.Error(w, errMsg, http.StatusNotImplemented)
}
