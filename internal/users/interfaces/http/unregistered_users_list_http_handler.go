package http

import (
	"github/grzegab/calendar/internal/users/application/unregistered_user_list"
	"net/http"
)

type UnregisteredUsersListHttpHandler struct {
	service *unregistered_user_list.Handler
}

func NewUnregisteredUsersListHttpHandler(svc *unregistered_user_list.Handler) *UnregisteredUsersListHttpHandler {
	return &UnregisteredUsersListHttpHandler{service: svc}
}

func (h *UnregisteredUsersListHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not initialized", http.StatusNotImplemented)
}
