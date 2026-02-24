package http

import (
	"github/grzegab/calendar/internal/scheduling/application/confirm_schedule"
	"net/http"
)

type ConfirmScheduleHttpHandler struct {
	service *confirm_schedule.Handler
}

func CreateConfirmScheduleHttpHandler(svc *confirm_schedule.Handler) *ConfirmScheduleHttpHandler {
	return &ConfirmScheduleHttpHandler{service: svc}
}

func (h *ConfirmScheduleHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {

}
