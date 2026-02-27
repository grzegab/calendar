package http

import (
	"github/grzegab/calendar/internal/users/application/user_details"
	"github/grzegab/calendar/internal/users/domain"
	"net/http"
)

type UserDetailsHttpHandler struct {
	service *user_details.Handler
}

func NewUserDetailsHttpHandler(service *user_details.Handler) *UserDetailsHttpHandler {
	return &UserDetailsHttpHandler{service: service}
}

func (h *UserDetailsHttpHandler) Handle(w http.ResponseWriter, r *http.Request) (*domain.User, error) {
	return nil, nil
}
