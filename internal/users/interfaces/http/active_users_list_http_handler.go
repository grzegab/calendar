package http

import (
	"github/grzegab/calendar/internal/users/application/active_user_list"
	"net/http"
)

type ActiveUsersHttpHandler struct {
	service *active_user_list.Handler
}

func NewActiveUsersHandler(svc *active_user_list.Handler) *ActiveUsersHttpHandler {
	return &ActiveUsersHttpHandler{service: svc}
}

func (h *ActiveUsersHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {}
