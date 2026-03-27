package http

import (
	"fmt"
	"github/grzegab/calendar/internal/scheduling/application/schedule_list"
	"github/grzegab/calendar/internal/shared/infrastructure/auth"
	"net/http"
)

type ListSlotsHttpHandler struct {
	service *schedule_list.Handler
}

func NewListSlotsHttpHandler(svc *schedule_list.Handler) *ListSlotsHttpHandler {
	return &ListSlotsHttpHandler{service: svc}
}

func (h *ListSlotsHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	teacherID, ok := auth.UserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	errMsg := fmt.Sprintf("not implemented for teacher %s le", teacherID)
	http.Error(w, errMsg, http.StatusNotImplemented)
}
