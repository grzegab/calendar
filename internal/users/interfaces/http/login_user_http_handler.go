package http

import (
	"github/grzegab/calendar/internal/users/application/login_user"
	"net/http"
)

type LoginUserHttpHandler struct {
	svc *login_user.Handler
}

func NewLoginHttpHandler(svc *login_user.Handler) *LoginUserHttpHandler {
	return &LoginUserHttpHandler{svc: svc}
}

func (h *LoginUserHttpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// @TODO: send SMS or Email on user
}
