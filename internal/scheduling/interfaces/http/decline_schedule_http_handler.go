package http

import (
	"fmt"
	"github/grzegab/calendar/internal/scheduling/application/decline_schedule"
	"github/grzegab/calendar/internal/shared/infrastructure/auth"
	"net/http"
)

type DeclineScheduleHttpHandler struct {
	service *decline_schedule.Handler
}

func CreateDeclineScheduleHttpHandler(svc *decline_schedule.Handler) *DeclineScheduleHttpHandler {
	return &DeclineScheduleHttpHandler{service: svc}
}

func (h *DeclineScheduleHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	teacherID, ok := auth.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	errMsg := fmt.Sprintf("not implemented for teacher %s le", teacherID)
	http.Error(w, errMsg, http.StatusNotImplemented)
}
