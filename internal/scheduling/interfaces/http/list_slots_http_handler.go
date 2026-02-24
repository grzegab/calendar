package http

import (
	"github/grzegab/calendar/internal/scheduling/application/schedule_list"
	"net/http"
)

type ListSlotsHttpHandler struct {
	service *schedule_list.Handler
}

func NewListSlotsHttpHandler(svc *schedule_list.Handler) *ListSlotsHttpHandler {
	return &ListSlotsHttpHandler{service: svc}
}

func (h *ListSlotsHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {}
