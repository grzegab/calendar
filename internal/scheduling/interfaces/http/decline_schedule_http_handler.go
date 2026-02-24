package http

import (
	"github/grzegab/calendar/internal/scheduling/application/decline_schedule"
	"net/http"
)

type DeclineScheduleHttpHandler struct {
	service *decline_schedule.Handler
}

func CreateDeclineScheduleHttpHandler(svc *decline_schedule.Handler) *DeclineScheduleHttpHandler {
	return &DeclineScheduleHttpHandler{service: svc}
}

func (h *DeclineScheduleHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {

}
