package http

import (
	"github/grzegab/calendar/internal/users/application/registered_user_list"
	"net/http"
)

type RegisteredUsersListHttpHandler struct {
	service *registered_user_list.Handler
}

func NewRegisteredUsersListHttpHandler(svc *registered_user_list.Handler) *RegisteredUsersListHttpHandler {
	return &RegisteredUsersListHttpHandler{service: svc}
}

func (h *RegisteredUsersListHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not initialized", http.StatusNotImplemented)
}
